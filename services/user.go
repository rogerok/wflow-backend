package services

import (
	"fmt"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/repositories"
	"golang.org/x/crypto/bcrypt"
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

func (s *userService) CreateUser(user *forms.UserCreateForm) (*string, error) {
	exists, err := s.r.CheckEmailExists(user.Email)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("%s", errors_utils.ErrEmailAlreadyExists)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
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
