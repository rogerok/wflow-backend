package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Age          int            `json:"age"`
	CreatedAt    time.Time      `json:"created_at"`
	Email        string         `json:"email"`
	FirstName    string         `json:"first_name"`
	Id           uuid.UUID      `json:"id"`
	LastName     sql.NullString `json:"last_name"`
	MiddleName   sql.NullString `json:"middle_name"`
	TelegramName sql.NullString `json:"telegram_name"`
	UpdatedAt    time.Time      `json:"updated_at"`
}
