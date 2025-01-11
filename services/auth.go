package services

import (
	"fmt"
	"github.com/rogerok/wflow-backend/errors_utils"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/responses"
	"github.com/rogerok/wflow-backend/utils"
)

type AuthService interface {
	Auth(user *forms.AuthForm) (resp *responses.TokensModel, err error)
	Refresh(rt string) (resp *responses.TokensModel, err error)
}

type authService struct {
	userRepo repositories.UserRepository
	authRepo repositories.AuthRepository
}

func NewAuthService(userRepo repositories.UserRepository, authRepo repositories.AuthRepository) AuthService {
	return &authService{userRepo: userRepo, authRepo: authRepo}
}

func (s *authService) Auth(loginForm *forms.AuthForm) (resp *responses.TokensModel, err error) {

	userData, err := s.userRepo.UserByEmail(loginForm.Email)

	if err != nil {
		return nil, errors_utils.CreateErrorMsg(errors_utils.ErrEmailOrPasswordError)
	}

	if !utils.ComparePassword(userData.Password, loginForm.Password) {
		return nil, errors_utils.CreateErrorMsg(errors_utils.ErrEmailOrPasswordError)
	}

	tokens, err := utils.CreateTokenPair(userData.Id)

	if err != nil {
		return nil, err
	}

	err = s.authRepo.CreateSession(userData.Id, tokens.RefreshToken)

	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *authService) Refresh(rt string) (resp *responses.TokensModel, err error) {
	sessionData, err := s.authRepo.GetByRefreshToken(rt)

	if err != nil {
		return nil, errors_utils.CreateErrorMsg(errors_utils.RefreshTokenNotFound)
	}

	tokens, err := utils.CreateTokenPair(sessionData.UserId)

	if err != nil {
		return nil, err
	}

	err = s.authRepo.CreateSession(sessionData.UserId, tokens.RefreshToken)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", sessionData)

	return tokens, nil
}
