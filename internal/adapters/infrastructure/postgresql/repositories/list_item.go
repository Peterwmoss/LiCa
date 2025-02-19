package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql/mappers"
	"github.com/Peterwmoss/LiCa/internal/core"
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
		Where("? = ?", bun.Ident("list__user.email"), string(user.Email)).
		Relation("List").
		Relation("List.User").
		Relation("Category").
		Relation("Product").
		Relation("Product.Categories").
		Relation("Product.Categories.Category").
		Scan(ctx)
	if err != nil {
		return []domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.GetAll: failed to get items for list: %s for user: %v. Error: %w", listName, user, err)
	}

	items, err := mappers.Map(dbItems, mappers.DbListItemToDomain)
	if err != nil {
		return []domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.GetAll: Error: %w", err)
	}

	return items, nil
}

func (r *ListItemRepository) GetById(ctx context.Context, user domain.User, id uuid.UUID) (domain.ListItem, error) {
	var dbItem postgresql.ListItem

	err := r.db.NewSelect().
		Model(&dbItem).
		Where("? = ?", bun.Ident("list__user.email"), string(user.Email)).
		Where("? = ?", bun.Ident("li.id"), id).
		Relation("List").
		Relation("List.User").
		Relation("Category").
		Relation("Product").
		Relation("Product.Categories").
		Relation("Product.Categories.Category").
		Limit(1).
		Scan(ctx)
	if err != nil {
		err = fmt.Errorf("repositories.ListItemRepository.GetById: failed to get item with id: %s for user: %v. Error: %w", id, user, err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ListItem{}, fmt.Errorf("%w. Error: %w", core.ErrNotFound, err)
		}
		return domain.ListItem{}, err
	}

	item, err := mappers.DbListItemToDomain(dbItem)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.GetById: Error: %w", err)
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

		if err != nil {
			return fmt.Errorf("repositories.ListItemRepository.Create: failed to persist item: %v. Error: %w", dbItem, err)
		}

		return nil
	})
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.Create: Error: %w", err)
	}

	created, err := r.GetById(ctx, user, item.Id)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("repositories.ListItemRepository.Create: Error: %w", err)
	}

	return created, nil
}

func (r *ListItemRepository) Update(ctx context.Context, user domain.User, item domain.ListItem) (domain.ListItem, error) {
	return domain.ListItem{}, errors.New("repositories.ListItemRepository.Update: not implemented")
}

func (r *ListItemRepository) Remove(ctx context.Context, user domain.User, id uuid.UUID) error {
	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbItem := postgresql.ListItem{}

		found, err := r.GetById(ctx, user, id)
		if err != nil {
			return fmt.Errorf("repositories.ListItemRepository.Remove: Error: %w", err)
		}

		if found.Id == uuid.Nil {
			return fmt.Errorf("repositories.ListItemRepository.Remove: not found: %s. Error: %w", id, core.ErrNotFound)
		}

		_, err = tx.NewDelete().
			Model(&dbItem).
			Where("id = ?", id).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("repositories.ListItemRepository.Remove: failed to delete: %s:. Error: %w", id, err)
		}
		return nil
	})
}
