package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/services"
	"github.com/rogerok/wflow-backend/utils"
)

// CreateReport godoc
// @Summary CreateReport Report
// @Description CreateReport Report
// @Tags Report
// @Param request body forms.ReportCreateForm true "body"
// @Produce json
// @Success 200 {object} responses.CreateResponse
// @Router /private/reports [post]
func CreateReport(s services.Reports) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		formData := new(forms.ReportCreateForm)
		userId, err := utils.GetSubjectFromHeaderToken(ctx)

		if err != nil {
			return err
		}

		if err := ctx.BodyParser(formData); err != nil {
			return utils.GetBadRequestError(ctx, err.Error())
		}
		formData.UserId = userId

		id, err := s.Create(formData)

		if err != nil {
			return utils.GetBadRequestError(ctx, err.Error())

		}

		return utils.GetResponseCreate(ctx, id)

	}
}
