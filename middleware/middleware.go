package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/utils"
	"strings"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")

		if token == "" {
			return utils.GetUnauthorizedErr(ctx)
		}

		_, err := utils.ParseToken(strings.Split(token, "Bearer ")[1])

		if err != nil {
			return utils.GetUnauthorizedErr(ctx)
		}

		return ctx.Next()
	}
}
