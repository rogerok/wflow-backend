package forms

import (
	"github.com/go-playground/validator/v10"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms/validators"
	"time"
)

type Pseudonym struct {
	FirstName *string `json:"firstName" validate:"omitempty,min=2,max=50"`
	LastName  *string `json:"lastName" validate:"omitempty,min=2,max=50"`
}

type Social struct {
	Instagram *string `json:"instagram" validate:"omitempty,url"`
	Telegram  *string `json:"telegram" validate:"omitempty,url"`
	TikTok    *string `json:"tiktok" validate:"omitempty,url"`
	Vk        *string `json:"vk" validate:"omitempty,url"`
}

type UserCreateForm struct {
	Email           string     `json:"email" validate:"required,email,max=255"`
	FirstName       string     `json:"firstName" validate:"required,min=2,max=50"`
	LastName        *string    `json:"lastName" validate:"omitempty,min=2,max=50"`
	MiddleName      *string    `json:"middleName" validate:"omitempty,min=2,max=50"`
	Password        string     `json:"password" validate:"required,min=8,max=255,passwordValidator"`
	PasswordConfirm string     `json:"passwordConfirm" validate:"required,min=8,max=255,eqfield=Password"`
	Pseudonym       *Pseudonym `json:"pseudonym" validate:"required"`
	SocialLinks     *Social    `json:"socialLinks" validate:"required"`
	BornDate        *time.Time `json:"bornDate" validate:"omitempty,futureDateValidator"`
}

func (uf *UserCreateForm) Validate() error {

	validators.RegisterTranslator(validators.PasswordValidatorName, errors_utils.ErrInvalidPassword)
	validators.RegisterTranslator(validators.ForbidFutureDateValidatorName, errors_utils.ErrFutureDate)

	customValidators := []func(v *validator.Validate) error{
		validators.RegisterPasswordValidator,
		validators.RegisterForbidFutureDateValidator,
	}

	if err := validators.ValidateWithCustomValidator(uf, customValidators); err != nil {
		return err

	}

	return nil
}
