package middleware

import (
	"fmt"
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

		parsedToken, err := utils.ParseToken(strings.Split(token, "Bearer ")[1])

		if err != nil {
			return utils.GetBadRequestError(ctx, err)
		}

		userId := parsedToken["sub"]

		fmt.Printf("%v\n", userId)
		return ctx.Next()
	}
}
