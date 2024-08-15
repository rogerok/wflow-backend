package errors

const (
	PasswordValidationErr = "PASSWORD_VALIDATION_ERROR"
)

func PasswordValidationError(field string) *CustomError {
	return New(PasswordValidationErr, "Пароль должен содержать только латинские буквы, цифры и спец символы. Должны быть минимум 1 строчная, 1 заглавная, 1 цифра и 1 один спецсимвол", field)
}
