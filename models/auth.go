package models

import (
	"github.com/google/uuid"
	"time"
)

type AuthSession struct {
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	ExpiresAt    *time.Time `json:"expiresAt" db:"expires_at"`
	IsRevoked    bool       `json:"isRevoked" db:"is_revoked"`
	RefreshToken string     `json:"refreshToken" db:"refresh_token"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	UserId       uuid.UUID  `json:"userId" db:"user_id"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
