package forms

import (
	"fmt"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	rutranslations "github.com/go-playground/validator/v10/translations/ru"
	"log"
)

type ErrorMap map[string]string

var validate *validator.Validate = nil
var trans ut.Translator

func InitTranslation() {
	ruLocale := ru.New()
	uni := ut.New(ruLocale, ruLocale)

	trans, _ = uni.GetTranslator("ru")
	validate = validator.New()

	err := rutranslations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		log.Fatalf(err.Error())
	}

}

func GetValidator() *validator.Validate {
	return validate
}

func FormatValidationError(err error) error {
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
