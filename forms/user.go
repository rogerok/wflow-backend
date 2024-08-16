package forms

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
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
	Password    string     `json:"-" validate:"required,min=8,max=255,passwordValidator"`
	Pseudonym   Pseudonym  `json:"pseudonym" validate:"required"`
	SocialLinks Social     `json:"socialLinks" validate:"required"`
	BornDate    *time.Time `json:"bornDate" db:"omitempty,datetime=2006-01-02"`
}

func passwordValidator(fl validator.FieldLevel) (check bool) {
	checks := []string{
		"[0-9]", // Checks for at least one digit
		"[a-z]", // Checks for at least one lowercase letter
		"[A-Z]", // Checks for at least one uppercase letter
		"[!@#$%^&*()\\-\\+}{'\";:.,></\\?\\|_=`~]", // Checks for at least one special character
	}

	pattern := fmt.Sprintf(`^(%s)+$`, strings.Join(checks, "|"))

	checks = append(checks, pattern)

	for _, c := range checks {
		if check = regexp.MustCompile(c).MatchString(fl.Field().String()); !check {
			return false
		}
	}

	return true

}

//func passwordErrorFunc(err validator.FieldError) error {
//	return errors.PasswordValidationError(err.Namespace())
//}

func registerValidator(v *validator.Validate) {
	v.RegisterValidation("passwordValidator", passwordValidator)
}

func (uf *UserCreateForm) Validate() (err error) {
	v := validator.New()

	//errMap := errors.ErrorMap{}

	registerValidator(v)

	//errValidation := v.Struct(uf)

	//if errValidation != nil {
	//	var validationErrs validator.ValidationErrors
	//	if errors.Error().As(err, &validationErrs) {
	//		// Iterate over each field error
	//		for _, fieldErr := range validationErrs {
	//			fmt.Printf("Field: %s, Error: %s\n", fieldErr.Field(), fieldErr.Error())
	//		}
	//	}
	//
	//	return nil

	return err
}
