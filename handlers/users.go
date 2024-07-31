package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

// UsersList godoc
// @Summary Get users list
// @Description Get users list
// @Tags User
// @Produce json
// @Success 200 {object} []models.User
// @Router /user [get]
func UsersList(s services.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := s.UsersList(1, 10)
		if err != nil {
			return err
		}
		return ctx.Status(fiber.StatusOK).JSON(users)
	}
}

// UserById godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags User
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func UserById(s services.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("invalid id"))
		}

		user, err := s.UserById(id)

		if err != nil {

			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(user)
	}
}
