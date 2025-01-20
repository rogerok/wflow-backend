package forms

import (
	"github.com/go-playground/validator/v10"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms/validators"
	"time"
)

type GoalCreateForm struct {
	EndDate     time.Time `json:"endDate" validate:"required,gt=StartDate,pastDateValidator"`
	GoalWords   int       `json:"goalWords" validate:"required,min=2"`
	StartDate   time.Time `json:"startDate" validate:"required,pastDateValidator"`
	Title       string    `json:"title" validate:"required,min=2,max=255"`
	UserId      string    `json:"-"`
	Description *string   `json:"description" validate:"omitempty,min=2,max=255"`
	BookId      string    `json:"bookId" validate:"required,uuid4"`
}

func RegisterEndDateValidator(v *validator.Validate) error {
	if err := v.RegisterValidation("endDateValidator", validators.GoalEndDateValidator); err != nil {
		return err
	}

	return nil
}

func (gf *GoalCreateForm) Validate() error {

	customValidators := []func(v *validator.Validate) error{
		validators.RegisterForbidPastDateValidator,
		RegisterEndDateValidator,
	}

	if err := validators.ValidateWithCustomValidator(gf, customValidators); err != nil {
		return err
	}

	validators.RegisterTranslator("endDateValidator", errors_utils.ErrInvalidGoalEndDate)
	validators.RegisterTranslator(validators.ForbidPastDateValidatorName, errors_utils.ErrPastDate)

	return nil
}
