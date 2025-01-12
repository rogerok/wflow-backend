package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

func CreateBook(s services.BooksService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		formData := new(forms.BookCreateForm)

		if err := ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

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
