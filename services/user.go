package services

import (
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
)

type UserService interface {
	UsersList(page int, perPage int) (users *[]models.User, err error)
	UserById(id string) (user *models.User, err error)
	CreateUser(user *forms.UserCreateForm) (id *string, err error)
}

type userService struct {
	r repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{
		r: repository,
	}
}

func (s *userService) UsersList(page int, perPage int) (users *[]models.User, err error) {

	users, err = s.r.UsersList(page, perPage)

	return users, err

}

func (s *userService) UserById(id string) (user *models.User, err error) {
	user, err = s.r.UserById(id)
	return user, err
}

func (s *userService) CreateUser(user *forms.UserCreateForm) (id *string, err error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	userId := "123"

	return &userId, nil
}
