package validators

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var ForbidPastDateValidatorName = "pastDateValidator"
var ForbidFutureDateValidatorName = "futureDateValidator"
var NotBeforeTodayValidatorName = "notBeforeTodayValidator"

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

func NotBeforeTodayValidator(fl validator.FieldLevel) bool {
	fieldValue, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	// Get year, month, day components only
	fieldYear, fieldMonth, fieldDay := fieldValue.Date()
	nowYear, nowMonth, nowDay := time.Now().Date()

	// Create new dates with just the date components (no time)
	fieldDate := time.Date(fieldYear, fieldMonth, fieldDay, 0, 0, 0, 0, time.UTC)
	todayDate := time.Date(nowYear, nowMonth, nowDay, 0, 0, 0, 0, time.UTC)

	// Valid if fieldDate is on or after todayDate
	return !fieldDate.Before(todayDate)
}

func RegisterForbidPastDateValidator(v *validator.Validate) error {

	if err := v.RegisterValidation(ForbidPastDateValidatorName, ForbidPastDateValidator); err != nil {
		return err
	}

	return nil
}

func RegisterNotBeforeTodayValidator(v *validator.Validate) error {

	if err := v.RegisterValidation(NotBeforeTodayValidatorName, NotBeforeTodayValidator); err != nil {
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
