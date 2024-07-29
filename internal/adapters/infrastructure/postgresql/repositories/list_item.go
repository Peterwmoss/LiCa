package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql/mappers"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ListItemRepository struct {
	db *bun.DB
}

func NewListItemRepository(db *bun.DB) ports.ListItemRepository {
	return &ListItemRepository{
		db: db,
	}
}

func (r *ListItemRepository) GetAll(ctx context.Context, user domain.User, listName domain.ListName) ([]domain.ListItem, error) {
	var dbItems []postgresql.ListItem

	err := r.db.NewSelect().
		Model(&dbItems).
		Where("? = ?", bun.Ident("user.email"), string(user.Email)).
		Relation("List").
		Relation("List.User").
		Relation("Category").
		Relation("Product").
		Relation("Product.Categories").
		Relation("Product.Categories.Category").
		Scan(ctx)
	if err != nil {
		return []domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.GetAll: failed to get items for list: %s for user: %v\n%w", listName, user, err)
	}

	items, err := mappers.Map(dbItems, mappers.DbListItemToDomain)
	if err != nil {
		return []domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.GetAll: failed to map items\n%w", err)
	}

	return items, nil
}

func (r *ListItemRepository) GetById(ctx context.Context, user domain.User, id uuid.UUID) (domain.ListItem, error) {
	var dbItem postgresql.ListItem

	err := r.db.NewSelect().
		Model(&dbItem).
		Where("? = ?", bun.Ident("user.email"), string(user.Email)).
		Where("id = ?", id).
		Relation("List").
		Relation("List.User").
		Relation("Category").
		Relation("Product").
		Relation("Product.Categories").
		Relation("Product.Categories.Category").
		Scan(ctx)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.GetById: failed to get item with id: %s for user: %v\n%w", id, user, err)
	}

	item, err := mappers.DbListItemToDomain(dbItem)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.GetById: failed to map item\n%w", err)
	}

	return item, nil
}

func (r *ListItemRepository) Create(ctx context.Context, user domain.User, item domain.ListItem) (domain.ListItem, error) {
	err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbItem := postgresql.ListItem{
			Id:         item.Id,
			ListId:     item.List.Id,
			ProductId:  item.Product.Id,
			CategoryId: item.Category.Id,
			Unit:       string(item.Unit),
			Amount:     float32(item.Amount),
		}

		_, err := tx.NewInsert().
			Model(&dbItem).
			Exec(ctx)

		return fmt.Errorf("repositories.ListItemRepository.Create: failed to persit item: %v\n%w", dbItem, err)
	})
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("Lrepositories.istItemRepository.Create: failed to create item: %v\n%w", item, err)
	}

	return r.GetById(ctx, user, item.Id)
}

func (r *ListItemRepository) Update(ctx context.Context, user domain.User, item domain.ListItem) (domain.ListItem, error) {
  return domain.ListItem{}, errors.New("repositories.ListItemRepository.Update: not implemented")
}

func (r *ListItemRepository) Remove(ctx context.Context, user domain.User, id uuid.UUID) error {
	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbItem := postgresql.ListItem{}

		found, err := r.GetById(ctx, user, id)
		if err != nil {
			return fmt.Errorf("repositories.ListItemRepository.Remove: failed to get item: %s\n%w", id, err)
		}

		if found.Id == uuid.Nil {
			return fmt.Errorf("repositories.ListItemRepository.Remove: not found: %s", id)
		}

		_, err = tx.NewDelete().
			Model(&dbItem).
			Where("id = ?", id).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("repositories.ListItemRepository.Remove: failed to delete: %s:\n%w", id, err)
		}
		return nil
	})
}
