package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
)

type UserRepository interface {
	UsersList(page int, perPage int) (user *[]models.User, err error)
	UserById(id string) (user *models.User, err error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UsersList(page int, perPage int) (users *[]models.User, err error) {

	users = &[]models.User{}

	err = r.db.Select(users, "SELECT * FROM USERS")

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return users, nil
}

func (r *userRepository) UserById(id string) (user *models.User, err error) {

	user = &models.User{}

	err = r.db.Get(user, "SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user %s not found %w", id, err)
		} else {
			return nil, fmt.Errorf(err.Error())
		}
	}

	return user, nil
}
