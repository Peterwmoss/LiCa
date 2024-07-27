package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/google/uuid"
)

type ProductService interface {
	Create(ctx context.Context, name string) (domain.Product, error)
	Get(ctx context.Context, name string) (domain.Product, error)
	GetById(ctx context.Context, id uuid.UUID) (domain.Product, error)
	GetAllForUser(ctx context.Context) ([]domain.Product, error)
}

type ProductRepository interface {
	GetAll(ctx context.Context, user domain.User) ([]domain.Product, error)
	Get(ctx context.Context, user domain.User, name domain.ProductName) (domain.Product, error)
	GetById(ctx context.Context, user domain.User, id uuid.UUID) (domain.Product, error)
	Create(ctx context.Context, product domain.Product) error
	Update(ctx context.Context, product domain.Product) error
}
