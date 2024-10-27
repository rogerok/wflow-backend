package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	CreateSession(userId uuid.UUID, refreshToken string) error
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r authRepository) CreateSession(userId uuid.UUID, refreshToken string) error {

	query := `INSERT INTO sessions (user_id, refresh_token) VALUES ($1, $2)
				ON CONFLICT (user_id) DO UPDATE 
				SET refresh_token = excluded.refresh_token,
					updated_at = now(),
					is_revoked = false
`
	//query := `UPDATE sessions SET refresh_token = $1, user_id=$2, is_revoked = false WHERE user_id = $3`

	_, err := r.db.Exec(query, userId, refreshToken)

	fmt.Printf("%v", err)

	return err
}
