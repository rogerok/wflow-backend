package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

// UsersList godoc
// @Summary Get users list
// @Description Get users list
// @Tags User
// @Produce json
// @Success 200 {object} []models.User
// @Router /users [get]
func UsersList(s services.UsersService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		params := new(models.UserQueryParams)

		err := ctx.QueryParser(params)

		if err != nil {
			return utils.GetParamsParsingError(ctx)
		}

		users, err := s.UsersList(params)

		if err != nil {
			return utils.GetBadRequestError(ctx, err)
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
// @Router /users/{id} [get]
func UserById(s services.UsersService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		user, err := s.UserById(id)

		if err != nil {
			return utils.GetResponseError(ctx, errors_utils.New(fiber.StatusBadRequest, err.Error()))
		}

		return ctx.Status(fiber.StatusOK).JSON(user)
	}
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Param request body forms.UserCreateForm true "body"
// @Produce json
// @Success 200 {object} responses.CreateResponse
// @Router /users [post]
func CreateUser(s services.UsersService) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		formData := new(forms.UserCreateForm)
		if err := ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err)
		}

		if err := formData.Validate(); err != nil {
			return utils.GetBadRequestError(ctx, err)
		}

		id, err := s.CreateUser(formData)

		if err != nil {
			return utils.GetBadRequestError(ctx, err)
		}

		return utils.GetResponseCreate(ctx, id)
	}
}

// LoginUser Login user godoc
// @Summary Login User
// @Description Login User
// @Tags User
// @Param request body forms.LoginUser true "body"
// @Produce json
// @Success 200 {object} responses.LoginSuccess
// @Router /users/login [post]
func LoginUser(s services.UsersService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		formData := new(forms.UserLoginForm)

		if err := ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err)
		}

		strf := "adasad"

		return utils.GetResponseCreate(ctx, &strf)
	}
}
