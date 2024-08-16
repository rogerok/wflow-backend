package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/errors"
	"github.com/rogerok/wflow-backend/responses"
)

func GetResponseError(ctx *fiber.Ctx, err *errors.CustomError) error {
	return ctx.Status(err.StatusCode).JSON(err)
}

func GetResponseCreate(ctx *fiber.Ctx, id *string) error {
	return ctx.Status(fiber.StatusCreated).JSON(responses.CreateResponse{Id: id})
}
