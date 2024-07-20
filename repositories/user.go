package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
)

type UserRepository interface {
	UsersList(page int, perPage int) ([]models.User, error)
	UserById(id string) (user *models.User, err error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UsersList(page int, perPage int) (users []models.User, err error) {

	query := "SELECT * FROM users ORDER BY created_at DESC"
	err = r.db.Select(&users, query)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) UserById(id string) (user *models.User, err error) {

	user = &models.User{}

	query := "SELECT * FROM users WHERE id = $1"

	err = r.db.Get(user, query, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
