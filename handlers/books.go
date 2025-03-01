package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

// CreateBook godoc
// @Summary CreateBook Book
// @Description CreateBook Book
// @Tags Books
// @Param request body forms.BookCreateForm true "body"
// @Produce json
// @Success 200 {object} responses.CreateResponse
// @Router /private/books [post]
func CreateBook(s services.BooksService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		formData := new(forms.BookCreateForm)

		userId, err := utils.GetSubjectFromHeaderToken(ctx)

		if err != nil {
			return err
		}

		if err := ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		formData.UserId = userId

		if err := formData.Validate(); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		id, err := s.CreateBook(formData)

		if err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		return utils.GetResponseCreate(ctx, id)
	}
}

// GetBooksList Books list godoc
// @Summary GetBooksList
// @Description GetBooksList
// @Tags Books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Router /private/books/{id} [get]
func GetBooksList(s services.BooksService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := utils.GetSubjectUuidFromHeaderToken(ctx)

		if err != nil {
			return err
		}

		params := new(models.BooksQueryParams)

		err = ctx.QueryParser(params)

		if err != nil {
			return utils.GetParamsParsingError(ctx)
		}

		params.UserId = userId

		books, err := s.GetBooksByUserId(params)

		if err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(books)
	}
}

// GetBookById book  godoc
// @Summary Get book by id
// @Description Get book by id
// @Tags Books
// @Produce json
// @Param RequestBody body models.BooksQueryParams true "Query parameters for books list"
// @Success 200 {object} []models.Book
// @Router /private/books [get]
func GetBookById(s services.BooksService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := utils.GetSubjectUuidFromHeaderToken(ctx)

		if err != nil {
			return err
		}

		id := ctx.Params("id")

		book, err := s.GetBookById(id, userId.String())

		if err != nil {
			return utils.GetNotFoundError(ctx, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(book)

	}
}
