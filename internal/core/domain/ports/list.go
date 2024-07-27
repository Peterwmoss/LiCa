package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

type ListService interface {
	Create(ctx context.Context, name string) (domain.List, error)
	Get(ctx context.Context, name string) (domain.List, error)
	GetAllForUser(ctx context.Context) ([]domain.List, error)
}

type ListRepository interface {
	Get(ctx context.Context, user domain.User, name domain.ListName) (domain.List, error)
	GetAll(ctx context.Context, user domain.User) ([]domain.List, error)
	Create(ctx context.Context, list domain.List) error
	Update(ctx context.Context, list domain.List) error
}
