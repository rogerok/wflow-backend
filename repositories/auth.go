package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
)

type AuthRepository interface {
	CreateSession(userId uuid.UUID, refreshToken string) error
	GetByRefreshToken(refreshToken string) (session *models.AuthSession, err error)
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

	_, err := r.db.Exec(query, userId, refreshToken)

	fmt.Printf("%v", err)

	return err
}

func (r authRepository) GetByRefreshToken(refreshToken string) (session *models.AuthSession, err error) {

	session = &models.AuthSession{}

	query := `SELECT * FROM sessions WHERE refresh_token = $1`

	err = r.db.Get(session, query, refreshToken)

	return session, err
}