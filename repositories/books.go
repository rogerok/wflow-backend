package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/utils"
)

type BooksRepository interface {
	Create(book *forms.BookForm) (id *string, err error)
	Edit(book *forms.BookForm, bookId string) (status bool, err error)
	Delete(bookId string, userId string) (status bool, err error)
	GetById(id string, userId string) (book *models.Book, err error)
	GetListByUserId(params *models.BooksQueryParams) (book *[]models.Book, err error)
}

type booksRepository struct {
	db *sqlx.DB
}

func NewBooksRepository(db *sqlx.DB) BooksRepository {
	return &booksRepository{db: db}
}

func (r *booksRepository) Create(book *forms.BookForm) (id *string, err error) {

	query := `INSERT INTO books (user_id, book_name, description) VALUES ($1, $2, $3) RETURNING id`

	err = r.db.QueryRow(query, book.UserId, book.Name, book.Description).Scan(&id)

	return id, err
}

func (r *booksRepository) Edit(book *forms.BookForm, bookId string) (status bool, err error) {

	query := `UPDATE books SET book_name = $1, description = $2 WHERE id = $3 AND user_id=$4`

	_, err = r.db.Exec(query, book.Name, book.Description, bookId, book.UserId)

	if err != nil {
		fmt.Printf("Error updating book %v. %v", book.Name, err.Error())
		return false, err

	}

	return true, nil
}

func (r *booksRepository) Delete(bookId string, userId string) (status bool, err error) {

	query := `DELETE FROM books WHERE id = $1 AND user_id = $2`

	_, err = r.db.Exec(query, bookId, userId)

	if err != nil {
		fmt.Printf("Error deleting book %v. %v", bookId, err.Error())
		return false, err

	}

	return true, nil
}

func (r *booksRepository) GetById(id string, userId string) (book *models.Book, err error) {
	query := `SELECT * FROM books WHERE id=$1 AND user_id=$2`

	book = &models.Book{}

	err = r.db.Get(book, query, id, userId)

	if err != nil {
		return nil, errors_utils.GetDBNotFoundError("Book")
	}

	return book, nil
}

func (r *booksRepository) GetListByUserId(params *models.BooksQueryParams) (books *[]models.Book, err error) {
	offset, selectAll := utils.HandlePagination(params.Page, params.PerPage)

	query := `SELECT created_at, updated_at, description, id, book_name FROM books WHERE user_id=$1`
	query += utils.GetAllowedOrderBy(params.OrderBy)

	books = &[]models.Book{}

	if selectAll {
		err = r.db.Select(books, query, params.UserId)
	} else {
		query += utils.GetOffsetLimitQuery(params.PerPage, offset)

		err = r.db.Select(books, query, params.UserId)
	}

	if err != nil {
		return nil, err
	}

	return books, nil
}
