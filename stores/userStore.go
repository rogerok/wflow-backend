package stores

import "github.com/rogerok/wflow-backend/models"

type UserStore interface {
	GetUserById(string) (*models.User, error)
}
