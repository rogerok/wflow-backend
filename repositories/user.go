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
	CreateUser(user *models.User) (id string, err error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UsersList(page int, perPage int) (users *[]models.User, err error) {

	users = &[]models.User{}

	query := `
			SELECT id, email, created_at, updated_at, first_name, last_name, middle_name, born_date, password,
				json_build_object(
					'firstName', pseudonym_first_name,
					'lastName', pseudonym_last_name
				) AS pseudonym,
				json_build_object(
					'instagram', social_instagram,
					'telegram', social_telegram,
					'tiktok', social_tiktok,
					'vk', social_vk
				) AS "socialLinks"
			FROM users`

	err = r.db.Select(users, query)

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

func (r *userRepository) CreateUser(user *models.User) (id string, err error) {

	return "", nil
}
