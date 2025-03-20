package forms

import (
	"github.com/go-playground/validator/v10"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms/validators"
	"time"
)

type GoalCreateForm struct {
	BookId      string    `json:"bookId" validate:"required,uuid4"`
	Description *string   `json:"description" validate:"omitempty,min=2,max=255"`
	EndDate     time.Time `json:"endDate" validate:"required,gt=StartDate,pastDateValidator"`
	GoalWords   float64   `json:"goalWords" validate:"required,min=2"`
	StartDate   time.Time `json:"startDate" validate:"required,pastDateValidator"`
	Title       string    `json:"title" validate:"required,min=2,max=255"`
	UserId      string    `json:"-"`
}

type GoalEditForm struct {
	Description *string   `json:"description" validate:"omitempty,min=2,max=255"`
	EndDate     time.Time `json:"endDate" validate:"required,pastDateValidator"`
	GoalId      string    `json:"-"`
	GoalWords   float64   `json:"goalWords" validate:"required,min=2"`
	StartDate   time.Time `json:"startDate" validate:"required"`
	Title       string    `json:"title" validate:"required,min=2,max=255"`
	UserId      string    `json:"-"`
}

func RegisterEndDateValidator(v *validator.Validate) error {
	if err := v.RegisterValidation("endDateValidator", validators.GoalEndDateValidator); err != nil {
		return err
	}

	return nil
}

//func (gf *GoalCreateForm) Validate() error {
//
//	customValidators := []func(v *validator.Validate) error{
//		validators.RegisterForbidPastDateValidator,
//		RegisterEndDateValidator,
//	}
//
//	if err := validators.ValidateWithCustomValidator(gf, customValidators); err != nil {
//		return err
//	}
//
//	validators.RegisterTranslator("endDateValidator", errors_utils.ErrInvalidGoalEndDate)
//	validators.RegisterTranslator(validators.ForbidPastDateValidatorName, errors_utils.ErrPastDate)
//
//	return nil
//}

type Validatable interface {
	Validate() error
}

func validateForm(form Validatable) error {
	customValidators := []func(v *validator.Validate) error{
		validators.RegisterForbidPastDateValidator,
		RegisterEndDateValidator,
	}

	if err := validators.ValidateWithCustomValidator(form, customValidators); err != nil {
		return err
	}

	validators.RegisterTranslator("endDateValidator", errors_utils.ErrInvalidGoalEndDate)
	validators.RegisterTranslator(validators.ForbidPastDateValidatorName, errors_utils.ErrPastDate)

	return nil
}

func (gf *GoalCreateForm) Validate() error {
	return validateForm(gf)
}

func (gf *GoalEditForm) Validate() error {
	return validateForm(gf)
}
