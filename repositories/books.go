package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/utils"
)

type BooksRepository interface {
	Create(book *models.Book) (id *string, err error)
	GetById(id string) (book *models.Book, err error)
	GetByUserId(params *models.BooksQueryParams) (book *[]models.Book, err error)
}

type booksRepository struct {
	db *sqlx.DB
}

func NewBooksRepository(db *sqlx.DB) BooksRepository {
	return &booksRepository{db: db}
}

func (r *booksRepository) Create(book *models.Book) (id *string, err error) {

	query := `INSERT INTO books (user_id, book_name, description) VALUES ($1, $2, $3) RETURNING id`

	err = r.db.QueryRow(query, book.UserId, book.Name, book.Description).Scan(&id)

	return id, err
}

func (r *booksRepository) GetById(id string) (book *models.Book, err error) {
	query := `SELECT * FROM books WHERE id=$1`

	err = r.db.Get(&book, query, id)

	if err != nil {
		return nil, errors_utils.GetDBNotFoundError("Book")
	}

	return book, nil
}

func (r *booksRepository) GetByUserId(params *models.BooksQueryParams) (books *[]models.Book, err error) {

	offset, selectAll := utils.HandlePagination(params.Page, params.PerPage)
	orderBy := utils.GetAllowedOrderBy(params.OrderBy)
	books = &[]models.Book{}

	query := `SELECT * FROM books WHERE user_id=$1 ORDER BY $2`

	if selectAll {
		err = r.db.Select(books, query, params.UserId, orderBy)
	} else {
		query = query + ` LIMIT $3 OFFSET $4`
		err = r.db.Select(books, query, params.UserId, utils.GetAllowedOrderBy(params.OrderBy), params.PerPage, offset)
	}
	if err != nil {
		return nil, err
	}

	return books, nil
}
