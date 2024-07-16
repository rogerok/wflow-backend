package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/services"
)

func UsersList(s services.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := s.UsersList(1, 10)
		if err != nil {
			return err
		}
		return ctx.Status(fiber.StatusOK).JSON(users)
	}
}

func UserById(s services.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		fmt.Printf(ctx.Params("id"))

		users, err := s.UserById(ctx.Params("id"))

		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(users)
	}
}
