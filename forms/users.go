package forms

import (
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
	BornDate        *time.Time `json:"bornDate" db:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

func (uf *UserCreateForm) Validate() error {
	RegisterTranslator(validators.PasswordValidatorName, errors_utils.ErrInvalidPassword)

	if err := ValidateWithCustomValidator(uf, validators.RegisterPasswordValidator); err != nil {
		return err
	}

	return nil
}
