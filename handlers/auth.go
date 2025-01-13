package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
	"net/http"
)

// AuthUser Auth user godoc
// @Summary Auth User
// @Description Auth User
// @Tags Auth
// @Param request body forms.AuthForm true "body"
// @Produce json
// @Success 200 {object} responses.TokenResponse
// @Router /api/auth [post]
func AuthUser(s services.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		formData := new(forms.AuthForm)

		if err := ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		if err := formData.Validate(); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		tokens, err := s.Auth(formData)

		if err != nil {
			return utils.GetResponseError(ctx, errors_utils.New(fiber.StatusUnauthorized, err.Error()))
		}

		cookies := fiber.Cookie{
			Name:     "rt",
			Value:    tokens.RefreshToken,
			Expires:  utils.GetRefreshTokenExpTime(),
			Secure:   true,
			HTTPOnly: true,
		}

		ctx.Cookie(&cookies)

		return ctx.Status(http.StatusOK).JSON(models.AuthResponse{Token: tokens.Token})
	}
}

// Refresh  user godoc
// @Summary Refresh User token
// @Description Refresh User token
// @Tags Auth
// @Param request body nil false "body"
// @Produce json
// @Success 200 {object} responses.TokenResponse
// @Router /api/auth/refresh [post]
func Refresh(s services.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		rt := ctx.Cookies("rt")

		if rt == "" {
			return utils.GetUnauthorizedErr(ctx)
		}

		tokens, err := s.Refresh(rt)

		if err != nil {
			return utils.GetUnauthorizedErr(ctx)
		}

		cookies := fiber.Cookie{
			Name:     "rt",
			Value:    tokens.RefreshToken,
			Expires:  utils.GetRefreshTokenExpTime(),
			Secure:   true,
			HTTPOnly: true,
		}

		ctx.Cookie(&cookies)

		return ctx.Status(http.StatusOK).JSON(models.AuthResponse{Token: tokens.Token})
	}
}

func Logout(s services.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		rt := ctx.Cookies("rt")

		if rt == "" {
			return utils.GetBadRequestError(ctx, errors_utils.ErrRefreshTokenNotFound)
		}

		_, err := utils.ParseToken(rt)

		if err != nil {
			return utils.GetBadRequestError(ctx, errors_utils.ErrRefreshTokenNotFound)
		}

		return ctx.SendStatus(http.StatusOK)
	}
}
