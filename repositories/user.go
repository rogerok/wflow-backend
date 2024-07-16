package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/models"
)

type UserRepository interface {
	UsersList(page int, perPage int) ([]models.User, error)
	UserById(id string) (user models.User, err error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UsersList(page int, perPage int) (users []models.User, err error) {
	// Implementation of GetUserById method
	// Example:
	// query := "SELECT * FROM users WHERE id = $1"
	// var user models.User
	// err := r.db.Get(&user, query, userID)
	// return &user, err
	return nil, nil
}

func (r *userRepository) UserById(id string) (user models.User, err error) {
	return models.User{
		Age:          0,
		CreatedAt:    "123",
		Email:        "123",
		FirstName:    "123",
		LastName:     "123",
		MiddleName:   "123",
		TelegramName: "123",
		UpdatedAt:    "123",
	}, nil
}
