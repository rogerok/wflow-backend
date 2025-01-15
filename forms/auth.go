package forms

import (
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms/validators"
)

type AuthForm struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255,passwordValidator"`
}

func (f *AuthForm) Validate() error {
	RegisterTranslator("passwordValidator", errors_utils.ErrInvalidPassword)

	if err := ValidateWithCustomValidator(f, validators.RegisterPasswordValidator); err != nil {
		return err
	}

	return nil
}
