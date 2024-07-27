package services

import (
	"context"
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
		return domain.Category{}, err
	}

	category := domain.CreateCategory(domainName, user)

	err = s.repository.Create(ctx, category)
	if err != nil {
		return domain.Category{}, err
	}

	return s.Get(ctx, name)
}

// Get implements ports.CategoryService.
func (s *CategoryService) Get(ctx context.Context, name string) (domain.Category, error) {
	user := ctx.Value("user").(domain.User)

	domainName, err := domain.NewCategoryName(name)
	if err != nil {
		return domain.Category{}, err
	}

	return s.repository.Get(ctx, user, domainName)
}

// GetAllForUser implements ports.CategoryService.
func (s *CategoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	user := ctx.Value("user").(domain.User)

	return s.repository.GetAll(ctx, user)
}

// GetById implements ports.CategoryService.
func (s *CategoryService) GetById(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	user := ctx.Value("user").(domain.User)

	return s.repository.GetById(ctx, user, id)
}
