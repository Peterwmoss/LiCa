package services

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type ListService struct {
	repository     ports.ListRepository
	userRepository ports.UserRepository
}

func NewListService(repository ports.ListRepository, userRepository ports.UserRepository) ports.ListService {
	return &ListService{
		repository,
		userRepository,
	}
}

func (s *ListService) Create(ctx context.Context, name string) (domain.List, error) {
	domainName, err := domain.NewListName(name)
	if err != nil {
		return domain.List{}, err
	}

  user := ctx.Value("user").(domain.User)
	domainEmail, err := domain.NewEmail(string(user.Email))
	if err != nil {
		return domain.List{}, err
	}

	domainUser, err := s.userRepository.GetByEmail(ctx, domainEmail)
	if err != nil {
		return domain.List{}, err
	}

	domainList := domain.CreateList(domainName, domainUser)

	err = s.repository.Create(ctx, domainList)
	if err != nil {
		return domain.List{}, err
	}

	return domainList, nil
}

func (s *ListService) Get(ctx context.Context, name string) (domain.List, error) {
	domainName, err := domain.NewListName(name)
	if err != nil {
		return domain.List{}, err
	}

  user := ctx.Value("user").(domain.User)
	domainEmail, err := domain.NewEmail(string(user.Email))
	if err != nil {
		return domain.List{}, err
	}

	return s.repository.Get(ctx, domainEmail, domainName)
}

func (s *ListService) GetAllForUser(ctx context.Context) ([]domain.List, error) {
  user := ctx.Value("user").(domain.User)
	domainEmail, err := domain.NewEmail(string(user.Email))
	if err != nil {
		return nil, err
	}

	return s.repository.GetAllByEmail(ctx, domainEmail)
}
