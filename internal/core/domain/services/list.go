package services

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type ListService struct {
	repository ports.ListRepository
}

func NewListService(repository ports.ListRepository) ports.ListService {
	return &ListService{
		repository,
	}
}

func (s *ListService) Create(ctx context.Context, name string) (domain.List, error) {
	domainName, err := domain.NewListName(name)
	if err != nil {
		return domain.List{}, err
	}

	user := ctx.Value("user").(domain.User)

	domainList := domain.CreateList(domainName, user)

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

	return s.repository.Get(ctx, user, domainName)
}

func (s *ListService) GetAllForUser(ctx context.Context) ([]domain.List, error) {
	user := ctx.Value("user").(domain.User)

	return s.repository.GetAll(ctx, user)
}
