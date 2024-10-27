package models

import (
	"github.com/google/uuid"
	"time"
)

type AuthSession struct {
	UserId       uuid.UUID `json:"userId" db:"user_id"`
	IsRevoked    bool      `json:"isRevoked" db:"is_revoked"`
	ExpiresAt    time.Time `json:"expiresAt" db:"expires_at"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	RefreshToken string    `json:"refreshToken" db:"refresh_token"`
}
