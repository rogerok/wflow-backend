package utils

import "github.com/gofiber/fiber/v2"

func GetResponseError(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(fiber.Map{"error": err.Error()})
}
