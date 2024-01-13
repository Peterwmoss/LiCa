package domain

import (
	"github.com/Peterwmoss/LiCa/internal/functions/auth"
	"github.com/rs/zerolog/log"
)

type (
	User struct {
		Email   string
		Picture string
	}

	UserService interface {
		Get(string) (User, error)
	}

	userService struct{}
)

func NewUserService() UserService {
	return &userService{}
}

func (u userService) Get(token string) (User, error) {
	userInfo, err := auth.GetUserInfo(token)
	if err != nil {
		return User{}, err
	}
	log.Info().Msgf("User: %s", userInfo.Email)

	return User{
		Email:   userInfo.Email,
		Picture: userInfo.Picture,
	}, nil
}
