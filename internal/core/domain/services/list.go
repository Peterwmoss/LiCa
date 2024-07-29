package services

import (
	"context"
	"fmt"

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
		return domain.List{}, fmt.Errorf("services.ListService.Create: failed to create list name: %s\n%w", name, err)
	}

	user := ctx.Value("user").(domain.User)

	domainList := domain.CreateList(domainName, user)

	err = s.repository.Create(ctx, domainList)
	if err != nil {
		return domain.List{}, fmt.Errorf("services.ListService.Create: failed to create list: %v\n%w", domainList, err)
	}

	return domainList, nil
}

func (s *ListService) Get(ctx context.Context, name string) (domain.List, error) {
	domainName, err := domain.NewListName(name)
	if err != nil {
		return domain.List{}, fmt.Errorf("services.ListService.Get: failed to create list name: %s\n%w", domainName, err)
	}

	user := ctx.Value("user").(domain.User)

	list, err := s.repository.Get(ctx, user, domainName)
	if err != nil {
		return domain.List{}, fmt.Errorf("services.ListService.Get: failed to get list with name: %s for user %v\n%w", domainName, user, err)
	}

	return list, nil
}

func (s *ListService) GetAllForUser(ctx context.Context) ([]domain.List, error) {
	user := ctx.Value("user").(domain.User)

	lists, err := s.repository.GetAll(ctx, user)
	if err != nil {
		return []domain.List{}, fmt.Errorf("services.ListService.GetAllForUser: failed to get lists for user %v\n%w", user, err)
	}

	return lists, nil
}
