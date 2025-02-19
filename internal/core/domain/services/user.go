package services

import (
	"context"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

// UserService is the only service to not have the user in the context
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
		return domain.User{}, fmt.Errorf("services.UserService.Create: Error: %w", err)
	}

	user := domain.CreateUser(domainEmail)
	err = s.repository.Create(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("services.UserService.Create: Error: %w", err)
	}

	created, err := s.repository.GetByEmail(ctx, domainEmail)
	if err != nil {
		return domain.User{}, fmt.Errorf("services.UserService.Create: Error: %w", err)
	}

	return created, nil
}

func (s *UserService) Get(ctx context.Context, email string) (domain.User, error) {
	domainEmail, err := domain.NewEmail(email)
	if err != nil {
		return domain.User{}, fmt.Errorf("services.UserService.Get: Error: %w", err)
	}

	user, err := s.repository.GetByEmail(ctx, domainEmail)
	if err != nil {
		return domain.User{}, fmt.Errorf("services.UserService.Get: Error: %w", err)
	}

	return user, nil
}
