package services

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type UserService struct {
	repository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) ports.UserService {
	return &UserService{
		repository,
	}
}

func (s *UserService) Create(ctx context.Context, email string) (domain.User, error) {
	domainEmail, err := domain.NewEmail(email)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.CreateUser(domainEmail)
	err = s.repository.Create(ctx, user)

	return user, err
}

func (s *UserService) Get(ctx context.Context, email string) (domain.User, error) {
	domainEmail, err := domain.NewEmail(email)
	if err != nil {
		return domain.User{}, err
	}

	return s.repository.GetByEmail(ctx, domainEmail)
}
