package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var PasswordValidatorName = "passwordValidator"

func PasswordValidator(fl validator.FieldLevel) (check bool) {

	patterns := []string{
		`[0-9]`,
		`[a-z]`,
		`[A-Z]`,
		`[!@#$%^&*()\-+}{'"[:;>.?/_~\|]`,
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

func RegisterPasswordValidator(v *validator.Validate) error {
	if err := v.RegisterValidation(PasswordValidatorName, PasswordValidator); err != nil {
		return err
	}

	return nil
}
