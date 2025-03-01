package forms

import (
	"github.com/go-playground/validator/v10"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms/validators"
)

type AuthForm struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255,passwordValidator"`
}

func (f *AuthForm) Validate() error {
	validators.RegisterTranslator("passwordValidator", errors_utils.ErrInvalidPassword)

	customValidators := []func(v *validator.Validate) error{
		validators.RegisterPasswordValidator,
	}

	if err := validators.ValidateWithCustomValidator(f, customValidators); err != nil {
		return err
	}

	return nil
}
