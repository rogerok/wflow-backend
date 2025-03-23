package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

// CreateGoal godoc
// @Summary CreateGoal Goals
// @Description Create goal for book Goals
// @Tags Goals
// @Param request body forms.GoalCreateForm true "body"
// Produce json
// @Success 200 {object} responses.CreateResponse
// @Router /private/goals [post]
func CreateGoal(s services.GoalsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		formData := new(forms.GoalCreateForm)

		userId, err := utils.GetSubjectFromHeaderToken(ctx)

		if err != nil {
			return utils.GetInvalidTokenError(ctx)
		}

		if err = ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		formData.UserId = userId

		if err = formData.Validate(); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		id, err := s.Create(formData)

		if err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		return utils.GetResponseCreate(ctx, id)
	}
}

// EditGoal  godoc
// @Summary EditGoal Goals
// @Description EditGoal goal Goals
// @Tags Goals
// @Param request body forms.GoalEditForm true "body"
// Produce json
// @Success 200 {object} models.GoalUpdateResponse
// @Router /private/goals/edit/{id} [put]
func EditGoal(s services.GoalsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		formData := new(forms.GoalEditForm)

		userId, err := utils.GetSubjectFromHeaderToken(ctx)

		if err != nil {
			return utils.GetInvalidTokenError(ctx)
		}

		if err = ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		formData.UserId = userId
		formData.GoalId = ctx.Params("id")

		if err = formData.Validate(); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		data, err := s.Edit(formData)

		if err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(data)
	}
}

// DeleteGoal godoc
// @Summary DeleteGoal by id
// @Description DeleteGoal Book
// @Tags Goals
// @Param id path string true "goal ID"
// @Produce json
// @Success 200 {object} responses.StatusResponse
// @Router /private/goals/delete [delete]
func DeleteGoal(s services.GoalsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		userId, err := utils.GetSubjectFromHeaderToken(ctx)

		if err != nil {
			return err
		}

		goalId := ctx.Params("id")
		status, err := s.Delete(goalId, userId)

		if err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		return utils.GetSuccessResponse(ctx, status)
	}
}

// GetList  Goals list godoc
// @Summary Get goals list by book id
// @Description Get goals list by book id
// @Tags Goals
// @Produce json
// @Param RequestBody body models.GoalsQueryParams true "Query parameters for goals list"
// @Success 200 {object} []models.Goals
// @Router /private/goals [get]
func GetList(s services.GoalsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := utils.GetSubjectUuidFromHeaderToken(ctx)

		if err != nil {
			return err
		}

		params := new(models.GoalsQueryParams)

		err = ctx.QueryParser(params)

		params.UserId = userId

		if err != nil {
			return utils.GetParamsParsingError(ctx)
		}

		goals, err := s.GetList(params)

		if err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(goals)
	}
}

// GetGoalById  Goals list godoc
// @Summary Get by id
// @Description Get goal by id
// @Tags Goals
// @Produce json
// @Param id path string true "Goals id"
// @Success 200 {object} models.Goals
// @Router /private/goals/{id} [get]
func GetGoalById(s services.GoalsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		goal, err := s.GetById(id)

		if err != nil {
			return utils.GetResponseError(ctx, errors_utils.New(fiber.StatusBadRequest, err.Error()))
		}

		return ctx.Status(fiber.StatusOK).JSON(goal)
	}
}
