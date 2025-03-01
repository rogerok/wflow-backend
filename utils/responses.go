package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/responses"
)

func GetResponseError(ctx *fiber.Ctx, err *errors_utils.CustomError) error {
	return ctx.Status(err.StatusCode).JSON(err)
}

func GetResponseCreate(ctx *fiber.Ctx, id *string) error {
	return ctx.Status(fiber.StatusCreated).JSON(responses.CreateResponse{Id: id})
}

func GetSuccessResponse(ctx *fiber.Ctx, isSuccess bool) error {
	return ctx.Status(fiber.StatusOK).JSON(responses.StatusResponse{Status: isSuccess})
}

func GetBadRequestError(ctx *fiber.Ctx, err string) error {
	return GetResponseError(ctx, errors_utils.New(fiber.StatusBadRequest, err))
}

func GetParamsParsingError(ctx *fiber.Ctx) error {
	return GetResponseError(ctx, errors_utils.New(fiber.StatusBadRequest, errors_utils.ErrQueryParamsParsing))
}

func GetUnauthorizedErr(ctx *fiber.Ctx) error {
	return GetResponseError(ctx, errors_utils.New(fiber.StatusUnauthorized, errors_utils.ErrUnauthorized))
}

func GetNotFoundError(ctx *fiber.Ctx, err string) error {
	return GetResponseError(ctx, errors_utils.New(fiber.StatusNotFound, err))
}

func GetInvalidTokenError(ctx *fiber.Ctx) error {
	return GetResponseError(ctx, errors_utils.New(fiber.StatusUnauthorized, errors_utils.ErrInvalidToken))
}
