package errors_utils

import "fmt"

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func New(code int, message string) *CustomError {
	return &CustomError{
		StatusCode: code,
		Message:    message,
	}
}

func CreateErrorMsg(msgs ...string) error {
	combinedMsg := ""

	for _, msg := range msgs {
		combinedMsg += fmt.Sprintf("%s, ", msg)
	}

	return fmt.Errorf(combinedMsg)
}

func GetDBNotFoundError(entity string) error {
	return CreateErrorMsg(fmt.Sprintf("Cущность %s не найдена", entity))
}
