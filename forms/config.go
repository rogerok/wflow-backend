package forms

import (
	"fmt"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	rutranslations "github.com/go-playground/validator/v10/translations/ru"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rogerok/wflow-backend/utils"
)

var validate *validator.Validate = nil
var trans ut.Translator

func InitTranslation() {
	ruLocale := ru.New()
	uni := ut.New(ruLocale, ruLocale)

	trans, _ = uni.GetTranslator("ru")

	validate = validator.New(validator.WithRequiredStructEnabled())

	err := rutranslations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		log.Errorf(err.Error())
	}

}

func GetValidator() *validator.Validate {
	if validate == nil {
		log.Fatalf("Validator not initialized. Call InitTranslation first.")
	}

	return validate
}

func formatValidationError(err error) error {
	var errMsg string
	for _, err := range err.(validator.ValidationErrors) {

		errMsg += fmt.Sprintf("%s;", err.Translate(trans))

	}

	return fmt.Errorf(errMsg)
}

func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

func translate(ut ut.Translator, fe validator.FieldError) string {
	msg, err := ut.T(fe.Tag(), fe.Field(), fe.Param())

	if err != nil {
		return fe.Error()
	}

	return msg
}

func RegisterTranslator(tag string, msg string) {
	err := validate.RegisterTranslation(tag, trans, registerTranslator(tag, msg), translate)

	if err != nil {
		return
	}
}

func validateStruct(s interface{}, v *validator.Validate) error {
	if err := v.Struct(s); err != nil {
		return formatValidationError(err)
	}

	return nil
}

func ValidateStruct(s interface{}) error {
	v := GetValidator()

	return validateStruct(s, v)
}

func registerCustomValidator(v *validator.Validate, customValidator func(v *validator.Validate) error) error {
	if err := customValidator(v); err != nil {
		return err
	}

	return nil
}

func ValidateWithCustomValidator(s interface{}, customValidator func(v *validator.Validate) error) error {
	v := GetValidator()

	if err := registerCustomValidator(v, customValidator); err != nil {
		return err
	}

	return validateStruct(s, v)
}

func RegisterPasswordValidator(v *validator.Validate) error {
	if err := v.RegisterValidation("passwordValidator", utils.PasswordValidator); err != nil {
		return err
	}

	return nil
}
