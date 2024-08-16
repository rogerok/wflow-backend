package errors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ErrorMap map[string]string

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.StatusCode, e.Message)
}

func New(code int, message string) *CustomError {
	return &CustomError{
		StatusCode: code,
		Message:    message,
	}
}

//func (e *CustomError) Error() ErrorMap {
//	errMsg := make(ErrorMap)
//	errMsg[e.Field] = e.Message
//	return errMsg
//}

func (v ErrorMap) AddError(field, message string) {
	v[field] = message
}

func GetErrorsMap(v *validator.ValidationErrors) ErrorMap {
	var validationErrors ErrorMap

	return validationErrors
}
