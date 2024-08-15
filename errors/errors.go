package errors

import "fmt"

type CustomError struct {
	Code    string
	Message string
	Field   string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s: %s", e.Code, e.Message, e.Field)
}

func New(code, message, field string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Field:   field,
	}
}
