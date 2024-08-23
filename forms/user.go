package forms

import (
	"github.com/go-playground/validator/v10"
	"regexp"
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
	Email       string     `json:"email" validate:"required,email,max=255"`
	FirstName   string     `json:"firstName" validate:"required,min=2,max=50"`
	LastName    *string    `json:"lastName" validate:"omitempty,min=2,max=50"`
	MiddleName  *string    `json:"middleName" validate:"omitempty,min=2,max=50"`
	Password    string     `json:"password" validate:"required,min=8,max=255,passwordValidator"`
	Pseudonym   Pseudonym  `json:"pseudonym" validate:"required"`
	SocialLinks Social     `json:"socialLinks" validate:"required"`
	BornDate    *time.Time `json:"bornDate" db:"omitempty,datetime=2006-01-02"`
}

func passwordValidator(fl validator.FieldLevel) (check bool) {

	patterns := []string{
		`[0-9]`,                          // At least one digit
		`[a-z]`,                          // At least one lowercase letter
		`[A-Z]`,                          // At least one uppercase letter
		`[!@#$%^&*()\-+}{'"[:;>.?/_~\|]`, // At least one special character
	}

	password := fl.Field().String()

	for _, pattern := range patterns {
		match, _ := regexp.MatchString(pattern, password)
		if !match {
			return false
		}
	}

	return true

}

func (uf *UserCreateForm) Validate() error {
	v := GetValidator()

	v.RegisterValidation("passwordValidator", passwordValidator)

	if err := v.Struct(uf); err != nil {
		return FormatValidationError(err)
	}

	return nil
}
