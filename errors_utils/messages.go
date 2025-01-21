package errors_utils

const (
	ErrInvalidPassword      = "Пароль должен содержать только латинские буквы, цифры и спец символы. Должны быть минимум 1 строчная, 1 заглавная, 1 цифра и 1 один спецсимвол"
	ErrEmailAlreadyExists   = "Пользователь с таким email уже существует"
	ErrCheckingUnique       = "Ошибка проверки уникальности"
	ErrHashing              = "Ошибка хеширования"
	ErrQueryParamsParsing   = "Ошибка обработки параметров запроса"
	ErrEmailOrPasswordError = "Пользователя с таким email не существует, или введён неверный пароль"
	ErrInvalidToken         = "Неверный токен"
	ErrTokenExpired         = "Токен просрочен"
	ErrUnauthorized         = "Требуется авторизация"
	ErrRefreshTokenNotFound = "Refresh token не найден"
	ErrLogout               = "Ошибка разлогина"
	ErrInvalidGoalEndDate   = "Дата окончания цели не может быть ранее даты начала"
	ErrPastDate             = "Дата не может быть прошедшей"
	ErrFutureDate           = "Дата не может быть будущей"
)
