package user

import (
	"context"
	"github.com/Peterwmoss/LiCa/internal/core/domain/user"
	"github.com/Peterwmoss/LiCa/internal/core/ports/user"
)

type service struct {
	command ports.UserCommandRepository
	query   ports.UserQueryRepository
}

func New(command ports.UserCommandRepository, query ports.UserQueryRepository) ports.UserPort {
	return &service{
		command,
		query,
	}
}

func (s *service) Create(ctx context.Context, email string) (*user.User, error) {
	domainEmail, err := user.NewEmail(email)
	if err != nil {
		return nil, err
	}

	user := user.New(domainEmail)
	s.command.Create(ctx, user)

	return user, nil
}

func (s *service) Get(ctx context.Context, email string) (*user.User, error) {
	domainEmail, err := user.NewEmail(email)
	if err != nil {
		return nil, err
	}

	user, err := s.query.UserByEmail(ctx, domainEmail)
	if err != nil {
		return nil, err
	}

	return user, nil
}
