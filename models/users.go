package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	TelegramName *string   `json:"telegramName" db:"telegram_name"`
	FirstName    string    `json:"firstName" db:"first_name"`
	LastName     *string   `json:"lastName" db:"last_name"`
	MiddleName   *string   `json:"middleName" db:"middle_name"`
	Age          *int      `json:"age" db:"age"`
}
