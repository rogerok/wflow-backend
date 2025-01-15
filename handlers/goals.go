package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

// CreateGoal godoc
// @Summary CreateGoal Goals
// @Description Create goal for book Goal
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
			return err
		}

		if err := ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		formData.UserId = userId

		if err := formData.Validate(); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		id, err := s.CreateGoal(formData)

		if err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}

		return utils.GetResponseCreate(ctx, id)
	}
}
