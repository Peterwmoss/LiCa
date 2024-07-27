package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/google/uuid"
)

type ListItemCreate struct {
	ListName     string
	ProductName  string
	CategoryName string
	Amount       float32
}

type ListItemUpdate struct {
	Id           uuid.UUID
	CategoryName string
	Amount       float32
}

type ListItemService interface {
	GetAll(ctx context.Context, listName string) ([]domain.ListItem, error)
	Add(ctx context.Context, createItem ListItemCreate) (domain.ListItem, error)
	Remove(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, updateItem ListItemUpdate) error
}

type ListItemRepository interface {
	GetAll(ctx context.Context, user domain.User, listName domain.ListName) ([]domain.ListItem, error)
	GetById(ctx context.Context, user domain.User, id uuid.UUID) (domain.ListItem, error)

	Create(ctx context.Context, user domain.User, item domain.ListItem) (domain.ListItem, error)
	Update(ctx context.Context, user domain.User, item domain.ListItem) (domain.ListItem, error)
	Remove(ctx context.Context, user domain.User, id uuid.UUID) error
}
