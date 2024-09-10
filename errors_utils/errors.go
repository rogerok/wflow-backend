package errors_utils

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

//func (e *CustomError) Error() string {
//	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
//}

func New(code int, message string) *CustomError {
	return &CustomError{
		StatusCode: code,
		Message:    message,
	}
}
