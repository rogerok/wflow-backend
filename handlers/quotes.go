package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

// GetRandomQuote book  godoc
// @Summary Get random quote
// @Description Get random quote
// @Tags Quotes
// @Produce json
// @Success 200 {object} models.Quotes
// @Router /private/quotes [get]
func GetRandomQuote(s services.QuotesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		book, err := s.GetRandom()

		if err != nil {
			return utils.GetNotFoundError(ctx, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(book)

	}
}
