package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain/list"
	"github.com/google/uuid"
)

type ListQueryRepository interface{}

type ListCommandRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*list.List, error)
	Save(ctx context.Context, list *list.List) error
}
