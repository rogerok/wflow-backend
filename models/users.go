package models

import (
	"time"
)

type User struct {
	Age          *int      `json:"age" db:"age"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	Email        string    `json:"email" db:"email"`
	FirstName    string    `json:"firstName" db:"first_name"`
	Id           string    `json:"id" db:"id"`
	LastName     *string   `json:"lastName" db:"last_name"`
	MiddleName   *string   `json:"middleName" db:"middle_name"`
	TelegramName *string   `json:"telegramName" db:"telegram_name"`
	UpdatedAt    *string   `json:"updatedAt" db:"updated_at"`
}
