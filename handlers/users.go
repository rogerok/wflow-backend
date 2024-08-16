package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/errors"
	"github.com/rogerok/wflow-backend/forms"
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
			return utils.GetResponseError(ctx, errors.New(fiber.StatusBadRequest, err.Error()))
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

		// TODO: it's smell
		if id == "" {
			return utils.GetResponseError(ctx, errors.New(fiber.StatusBadRequest, "Неверный id пользователя"))
		}

		user, err := s.UserById(id)

		if err != nil {
			return utils.GetResponseError(ctx, errors.New(fiber.StatusBadRequest, err.Error()))
		}

		return ctx.Status(fiber.StatusOK).JSON(user)
	}
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Param request body models.User true "body"
// @Produce json
// @Success 200 {object} responses.CreateResponse
// @Router /user [post]
func CreateUser(s services.UserService) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		formData := new(forms.UserCreateForm)

		if err := ctx.BodyParser(formData); err != nil {
			return utils.GetResponseError(ctx, errors.New(fiber.StatusBadRequest, err.Error()))
		}

		id, err := s.CreateUser(formData)

		if err != nil {
			return err
		}

		return utils.GetResponseCreate(ctx, id)
	}
}
