package services

import (
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
)

type BooksService interface {
	CreateBook(book *forms.BookCreateForm) (id *string, err error)
	GetBookById(id string, userId string) (book *models.Book, err error)
	GetBooksByUserId(params *models.BooksQueryParams) (book *[]models.Book, err error)
}

type booksService struct {
	r repositories.BooksRepository
}

func NewBooksService(r repositories.BooksRepository) BooksService {
	return &booksService{
		r: r,
	}
}

func mapFormToBookModel(book *forms.BookCreateForm) *models.Book {
	return &models.Book{
		Description: book.Description,
		Name:        book.Name,
		UserId:      book.UserId,
	}
}

func (s *booksService) CreateBook(book *forms.BookCreateForm) (id *string, err error) {

	bookData := mapFormToBookModel(book)

	id, err = s.r.Create(bookData)

	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *booksService) GetBookById(id string, userId string) (book *models.Book, err error) {
	book, err = s.r.GetById(id, userId)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *booksService) GetBooksByUserId(params *models.BooksQueryParams) (books *[]models.Book, err error) {
	books, err = s.r.GetListByUserId(params)

	if err != nil {
		return nil, err
	}

	return books, nil
}
