package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/google/uuid"
)

type CategoryService interface {
	Create(ctx context.Context, name string) (domain.Category, error)
	Get(ctx context.Context, name string) (domain.Category, error)
	GetById(ctx context.Context, id uuid.UUID) (domain.Category, error)
	GetAll(ctx context.Context) ([]domain.Category, error)
}

type CategoryRepository interface {
	GetAll(ctx context.Context, user domain.User) ([]domain.Category, error)
	Get(ctx context.Context, user domain.User, name domain.CategoryName) (domain.Category, error)
	GetById(ctx context.Context, user domain.User, id uuid.UUID) (domain.Category, error)
	Create(ctx context.Context, category domain.Category) error
	Update(ctx context.Context, category domain.Category) error
}
