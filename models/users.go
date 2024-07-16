package models

type User struct {
	Age          int    `json:"age"`
	CreatedAt    string `json:"created_at"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	Id           string `json:"id"`
	LastName     string `json:"last_name"`
	MiddleName   string `json:"middle_name"`
	TelegramName string `json:"telegram_Name"`
	UpdatedAt    string `json:"updated_at"`
}
