package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/forms"
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
