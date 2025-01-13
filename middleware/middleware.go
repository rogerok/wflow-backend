package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/utils"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		token, err := utils.GetAccessTokenFromHeader(ctx)

		if err != nil {
			return err

		}

		_, err = utils.ParseToken(token)

		if err != nil {
			return utils.GetUnauthorizedErr(ctx)
		}

		return ctx.Next()
	}
}
