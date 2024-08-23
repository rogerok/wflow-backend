package forms

import (
	"fmt"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	rutranslations "github.com/go-playground/validator/v10/translations/ru"
	"log"
)

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
		//if customMsg, found := ValidationErrorMessages[err.Tag()]; found {
		//	errMsg += fmt.Sprintf("%s;", customMsg)
		//} else {
		//	errMsg += fmt.Sprintf("%s %s;", err.Field(), err.Tag())
		//}

		errMsg += fmt.Sprintf("%s;", err.Translate(trans))

	}

	return fmt.Errorf(errMsg)
}
