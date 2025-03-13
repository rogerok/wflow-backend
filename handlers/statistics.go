package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

// GetUserStatistics statistics  godoc
// @Summary Get statistics by user id
// @Description get user's activity statistics
// @Tags Statistics
// @Produce json
// @Success 200 {object} models.UserStatistics
// @Router /private/statistics/user [get]
func GetUserStatistics(s services.StatisticsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := utils.GetSubjectUuidFromHeaderToken(ctx)

		if err != nil {
			return utils.GetInvalidTokenError(ctx)
		}

		data, err := s.GetUserStatistics(userId)

		if err != nil {
			return utils.GetNotFoundError(ctx, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(data)

	}
}

// GetGoalStatistics statistics  godoc
// @Summary Get statistics by goal id
// @Description get goal's activity statistics
// @Tags Statistics
// @Produce json
// @Param id path string true "Goal ID"
// @Success 200 {object} models.GoalStatistics
// @Router /private/statistics/goal/{id} [get]
func GetGoalStatistics(s services.StatisticsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		data, err := s.GetGoalStatistics(id)

		if err != nil {
			return utils.GetNotFoundError(ctx, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(data)

	}
}

// GetFullProfileChartData statistics  godoc
// @Summary Get profile full chart data user id
// @Description  Get profile full chart data user id
// @Tags Statistics
// @Produce json
// @Success 200 {object} models.FullProfileChartData
// @Router /private/statistics/user/full [get]
func GetFullProfileChartData(s services.StatisticsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := utils.GetSubjectUuidFromHeaderToken(ctx)

		if err != nil {
			return utils.GetInvalidTokenError(ctx)
		}

		data, err := s.GetFullProfileChartData(userId)

		if err != nil {
			return utils.GetNotFoundError(ctx, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(data)

	}
}
