package forms

import "github.com/rogerok/wflow-backend/errors_utils"

type AuthnForm struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255,passwordValidator"`
}

func (f *AuthnForm) Validate() error {
	RegisterTranslator("passwordValidator", errors_utils.ErrInvalidPassword)

	if err := ValidateWithCustomValidator(f, RegisterPasswordValidator); err != nil {
		return err
	}

	return nil
}
