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
		Where("? = ?", bun.Ident("user.email"), string(user.Email)).
		Relation("List").
		Relation("List.User").
		Relation("Category").
		Relation("Product").
		Relation("Product.Categories").
		Relation("Product.Categories.Category").
		Scan(ctx)
	if err != nil {
		return []domain.ListItem{}, err
	}

	return mappers.Map(dbItems, mappers.DbListItemToDomain)
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
		return domain.ListItem{}, err
	}

	return mappers.DbListItemToDomain(dbItem)
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

		return err
	})
	if err != nil {
		return domain.ListItem{}, err
	}

	return r.GetById(ctx, user, item.Id)
}

func (r *ListItemRepository) Update(ctx context.Context, user domain.User, item domain.ListItem) (domain.ListItem, error) {
	return domain.ListItem{}, errors.New("not implemented")
}

func (r *ListItemRepository) Remove(ctx context.Context, user domain.User, id uuid.UUID) error {
	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbItem := postgresql.ListItem{}

		found, err := r.GetById(ctx, user, id)
		if err != nil {
			return err
		}

		if found.Id == uuid.Nil {
			return fmt.Errorf("%w: %s", core.ErrNotFound, id.String())
		}

		_, err = tx.NewDelete().
			Model(&dbItem).
			Where("id = ?", id).
			Exec(ctx)

		return err
	})
}
