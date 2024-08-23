package errors

import (
	"fmt"
)

type ErrorMap map[string]string

//var ValidationErrorMessages = ErrorMap{
//	"passwordValidator": "Пароль должен содержать только латинские буквы, цифры и спец символы. Должны быть минимум 1 строчная, 1 заглавная, 1 цифра и 1 один спецсимвол",
//}

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

func New(code int, message string) *CustomError {
	return &CustomError{
		StatusCode: code,
		Message:    message,
	}
}
