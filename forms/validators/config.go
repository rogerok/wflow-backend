package validators

import (
	"fmt"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	rutranslations "github.com/go-playground/validator/v10/translations/ru"
	"github.com/gofiber/fiber/v2/log"
)

var Validate *validator.Validate = nil
var trans ut.Translator

func InitTranslation() {
	ruLocale := ru.New()
	uni := ut.New(ruLocale, ruLocale)

	trans, _ = uni.GetTranslator("ru")

	Validate = validator.New(validator.WithRequiredStructEnabled())

	err := rutranslations.RegisterDefaultTranslations(Validate, trans)

	if err != nil {
		log.Errorf(err.Error())
	}
}

func GetValidator() *validator.Validate {
	if Validate == nil {
		log.Fatalf("Validator not initialized. Call InitTranslation first.")
	}

	return Validate
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
	err := Validate.RegisterTranslation(tag, trans, registerTranslator(tag, msg), translate)

	if err != nil {
		return
	}
}

func ValidateStruct(s interface{}, v *validator.Validate) error {
	if err := v.Struct(s); err != nil {
		return formatValidationError(err)
	}
	return nil
}

func registerCustomValidators(v *validator.Validate, customValidators []func(v *validator.Validate) error) error {
	for _, customValidator := range customValidators {
		if err := customValidator(v); err != nil {
			return err
		}
	}
	return nil
}

func ValidateWithCustomValidator(s interface{}, customValidators []func(v *validator.Validate) error) error {
	v := GetValidator()

	if err := registerCustomValidators(v, customValidators); err != nil {
		return err
	}

	return ValidateStruct(s, v)
}
