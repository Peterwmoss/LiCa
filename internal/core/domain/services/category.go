package services

import (
	"context"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
)

type CategoryService struct {
	repository ports.CategoryRepository
}

func NewCategoryService(repository ports.CategoryRepository) ports.CategoryService {
	return &CategoryService{
		repository,
	}
}

// Create implements ports.CategoryService.
func (s *CategoryService) Create(ctx context.Context, name string) (domain.Category, error) {
	user := ctx.Value("user").(domain.User)

	domainName, err := domain.NewCategoryName(name)
	if err != nil {
		return domain.Category{}, fmt.Errorf("services.CategoryService.Create: failed to create domain name: %s\n%w", name, err)
	}

	category := domain.CreateCategory(domainName, user)

	err = s.repository.Create(ctx, category)
	if err != nil {
		return domain.Category{}, fmt.Errorf("services.CategoryService.Create: failed to create category: %s for user: %v\n%w", name, user, err)
	}

	created, err := s.Get(ctx, name)
	if err != nil {
		return domain.Category{}, fmt.Errorf("services.CategoryService.Create: failed to get category with name: %s for user: %v\n%w", name, user, err)
	}

	return created, nil
}

// Get implements ports.CategoryService.
func (s *CategoryService) Get(ctx context.Context, name string) (domain.Category, error) {
	user := ctx.Value("user").(domain.User)

	domainName, err := domain.NewCategoryName(name)
	if err != nil {
		return domain.Category{}, fmt.Errorf("services.CategoryService.Get: failed to create domain name: %s\n%w", name, err)
	}

	category, err := s.repository.Get(ctx, user, domainName)
	if err != nil {
		return domain.Category{}, fmt.Errorf("services.CategoryService.Get: failed to get category with name: %s for user: %v\n%w", name, user, err)
	}

	return category, nil
}

// GetAllForUser implements ports.CategoryService.
func (s *CategoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	user := ctx.Value("user").(domain.User)

	categories, err := s.repository.GetAll(ctx, user)
	if err != nil {
		return []domain.Category{}, fmt.Errorf("services.CategoryService.GetAll: failed to get categories for user: %v\n%w", user, err)
	}

	return categories, nil
}

// GetById implements ports.CategoryService.
func (s *CategoryService) GetById(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	user := ctx.Value("user").(domain.User)

	category, err := s.repository.GetById(ctx, user, id)
	if err != nil {
		return domain.Category{}, fmt.Errorf("services.CategoryService.GetById: failed to get category with id: %s for user: %v\n%w", id, user, err)
	}

	return category, nil
}
