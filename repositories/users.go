package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/utils"
)

type UserRepository interface {
	UsersList(page int, perPage int) (user *[]models.User, err error)
	UserById(id string) (user *models.User, err error)
	CheckEmailExists(email string) (exists bool, err error)
	CreateUser(user *models.User) (id *string, err error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UsersList(page int, perPage int) (users *[]models.User, err error) {

	offset, selectAll := utils.HandlePagination(page, perPage)

	users = &[]models.User{}

	query := `
			SELECT id, email, created_at, updated_at, first_name, last_name, middle_name, born_date,
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

	if selectAll {
		err = r.db.Select(users, query)
	} else {
		query = query + ` ORDER BY created_At DESC LIMIT $1 OFFSET $2`
		err = r.db.Select(users, query, perPage, offset)
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) UserById(id string) (user *models.User, err error) {

	user = &models.User{}

	query := `
			SELECT id, email, created_at, updated_at, first_name, last_name, middle_name,
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
		return nil, errors_utils.GetDBNotFoundError("User")
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *models.User) (id *string, err error) {

	query := `INSERT INTO users (
                  	email, password, first_name, last_name, middle_name, born_date, pseudonym_first_name, pseudonym_last_name,
                   social_instagram, social_telegram, social_tiktok, social_vk)
				VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`

	err = r.db.QueryRow(query, user.Email, user.Password, user.FirstName, user.LastName, user.MiddleName, user.BornDate,
		user.Pseudonym.FirstName, user.Pseudonym.LastName, user.SocialLinks.Instagram, user.SocialLinks.Telegram, user.SocialLinks.TikTok, user.SocialLinks.Vk).Scan(&id)

	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *userRepository) CheckEmailExists(email string) (exists bool, err error) {

	err = r.db.Get(&exists,
		"SELECT EXISTS(SELECT * FROM users  WHERE email=$1)", email)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRepository) LoginUser(user *forms.UserLoginForm) error {

	var password string

	err := r.db.Get(password, `SELECT password FROM users WHERE email=$1`, user.Email)

	if err != nil {
		return errors_utils.CreateErrorMsg(errors_utils.ErrEmailOrPasswordError)
	}

	if utils.ComparePassword(password, user.Password) {

	}

	return err
}
