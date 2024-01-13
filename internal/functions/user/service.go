package user

import (
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/functions/auth"
	"github.com/rs/zerolog/log"
)

type (
  Service interface {
    Get(string) (domain.User, error)
  }

  service struct {
  }
)

func NewService() Service {
  return &service{}
}

func (u service) Get(token string) (domain.User, error) {
  userInfo, err := auth.GetUserInfo(token)
  if err != nil {
    return domain.User{}, err
  }
  log.Info().Msgf("User: %s", userInfo.Email)

  return domain.User{
    Email: userInfo.Email,
    Picture: userInfo.Picture,
  }, nil
}
