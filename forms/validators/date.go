package validators

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var ForbidPastDateValidatorName = "pastDateValidator"
var ForbidFutureDateValidatorName = "futureDateValidator"

func GoalEndDateValidator(fl validator.FieldLevel) (check bool) {
	startDate := fl.Parent().FieldByName("StartDate").Interface().(time.Time)

	endDate := fl.Field().Interface().(time.Time)

	return endDate.After(startDate)
}

func ForbidPastDateValidator(fl validator.FieldLevel) (check bool) {

	fieldValue, ok := fl.Field().Interface().(time.Time)

	if !ok {
		return false
	}

	return !fieldValue.Truncate(24 * time.Hour).Before(time.Now().Truncate(24 * time.Hour))

}

func RegisterForbidPastDateValidator(v *validator.Validate) error {

	if err := v.RegisterValidation(ForbidPastDateValidatorName, ForbidPastDateValidator); err != nil {
		return err
	}

	return nil
}

func ForbidFutureDateValidator(fl validator.FieldLevel) bool {
	format := "2006-01-02"

	fieldValue, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	fieldDate := fieldValue.Format(format)

	today := time.Now().Format(format)

	return fieldDate <= today
}

func RegisterForbidFutureDateValidator(v *validator.Validate) error {

	if err := v.RegisterValidation(ForbidFutureDateValidatorName, ForbidFutureDateValidator); err != nil {
		return err
	}

	return nil
}
