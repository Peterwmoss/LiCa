package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

type ListItemService interface {
	Add(ctx context.Context, name string) (domain.List, error)
	Remove(ctx context.Context, name string) (domain.List, error)
	Update(ctx context.Context) ([]domain.List, error)
}

type ListItemRepository interface {
	Get(ctx context.Context, email domain.Email, name domain.ListName) (domain.List, error)
	GetAllByEmail(ctx context.Context, email domain.Email) ([]domain.List, error)
	Create(ctx context.Context, list domain.List) error
	Update(ctx context.Context, list domain.List) error
}
