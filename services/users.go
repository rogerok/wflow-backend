package services

import (
	"fmt"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/utils"
)

type UsersService interface {
	UsersList(page int, perPage int) (users *[]models.User, err error)
	UserById(id string) (user *models.User, err error)
	CreateUser(user *forms.UserCreateForm) (id *string, err error)
}

type usersService struct {
	r repositories.UserRepository
}

func NewUsersService(repository repositories.UserRepository) UsersService {
	return &usersService{
		r: repository,
	}
}

func (s *usersService) UsersList(page int, perPage int) (users *[]models.User, err error) {

	users, err = s.r.UsersList(page, perPage)

	return users, err

}

func (s *usersService) UserById(id string) (user *models.User, err error) {
	user, err = s.r.UserById(id)
	return user, err
}

func (s *usersService) CreateUser(user *forms.UserCreateForm) (*string, error) {
	exists, err := s.r.CheckEmailExists(user.Email)

	if err != nil {
		return nil, errors_utils.CreateErrorMsg(err, errors_utils.ErrCheckingUnique, "email")
	}

	if exists {
		return nil, fmt.Errorf("%s", errors_utils.ErrEmailAlreadyExists)
	}

	encryptedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, errors_utils.CreateErrorMsg(err, errors_utils.ErrHashing)
	}

	userData := models.User{
		BornDate:   user.BornDate,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Password:   string(encryptedPassword),
		Pseudonym: models.Pseudonym{
			FirstName: user.Pseudonym.FirstName,
			LastName:  user.Pseudonym.LastName,
		},
		SocialLinks: models.Social{
			Instagram: user.SocialLinks.Instagram,
			Telegram:  user.SocialLinks.Telegram,
			TikTok:    user.SocialLinks.TikTok,
			Vk:        user.SocialLinks.Vk,
		},
	}

	id, err := s.r.CreateUser(&userData)

	if err != nil {
		return nil, err
	}

	return id, nil
}
