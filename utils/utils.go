package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/responses"
)

func GetResponseError(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(fiber.Map{"error": err.Error()})
}

func GetResponseCreate(ctx *fiber.Ctx, id *string) error {
	return ctx.Status(fiber.StatusCreated).JSON(responses.CreateResponse{Id: id})
}
