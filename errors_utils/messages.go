package errors_utils

const (
	ErrInvalidPassword      = "Пароль должен содержать только латинские буквы, цифры и спец символы. Должны быть минимум 1 строчная, 1 заглавная, 1 цифра и 1 один спецсимвол"
	ErrEmailAlreadyExists   = "Пользователь с таким email уже существует"
	ErrCheckingUnique       = "Ошибка проверки уникальности"
	ErrHashing              = "Ошибка хеширования"
	ErrQueryParamsParsing   = "Ошибка обработки параметров запроса"
	ErrEmailOrPasswordError = "Пользователя с таким email не существует, или введён неверный пароль"
	ErrTokenParse           = "Некорректный метод авторизации"
	ErrUnauthorized         = "Требуется авторизация"
	RefreshTokenNotFound    = "Refresh token не найден"
)
