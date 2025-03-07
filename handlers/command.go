package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/services"
)

func RecalculateGoals(s services.GoalsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		s.RecalculateGoals()
		return nil
	}
}
