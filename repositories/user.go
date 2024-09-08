package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/models"
)

type UserRepository interface {
	UsersList(page int, perPage int) (user *[]models.User, err error)
	UserById(id string) (user *models.User, err error)
	//CreateUser(user *models.User) (id string, err error)
	//CheckEmailExists(email string) (exists bool, err error)
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
			FROM users WHERE id=$1`

	err = r.db.Get(user, query, id)

	if err != nil {
		if notFoundError := errors_utils.CheckNotFoundError(err, "User"); notFoundError != nil {
			return nil, notFoundError
		} else {
			return nil, fmt.Errorf(err.Error())
		}
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *models.User) (id string, err error) {

	return "232323", nil
}

//func (r *userRepository) CheckEmailExists(email string) (bool, error) {
//	var email string;
//
//	err := r.db.Get(&email,
//		"SELECT FROM USERS WHERE ")
//}
