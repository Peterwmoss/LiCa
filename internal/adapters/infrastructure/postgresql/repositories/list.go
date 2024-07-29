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

var (
	ErrListNotFound = errors.New("list does not exist")
)

type ListRepository struct {
	db *bun.DB
}

func NewListRepository(db *bun.DB) ports.ListRepository {
	return &ListRepository{
		db,
	}
}

func (r *ListRepository) Get(ctx context.Context, user domain.User, name domain.ListName) (domain.List, error) {
	dbList := postgresql.List{}

	err := r.db.NewSelect().
		Model(&dbList).
		Where("name = ?", name).
		Where("? like ?", bun.Ident("user.email"), string(user.Email)).
		Relation("User").
		Relation("ListItems").
		Relation("ListItems.List").
		Relation("ListItems.List.User").
		Relation("ListItems.Product").
		Relation("ListItems.Category").
		Relation("ListItems.Product.Categories").
		Relation("ListItems.Product.Categories.Category").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return domain.List{}, fmt.Errorf("repositories.ListRepository.Get: failed to get list with name: %s for user: %v:\n%w", name, user, err)
	}

	list, err := mappers.DbListToDomain(dbList)
	if err != nil {
		return domain.List{}, fmt.Errorf("repositories.ListRepository.Get: failed to map list\n%w", err)
	}

	return list, nil
}

func (r *ListRepository) GetAll(ctx context.Context, user domain.User) ([]domain.List, error) {
	var dbLists []postgresql.List

	err := r.db.NewSelect().
		Model(&dbLists).
		Where("? like ?", bun.Ident("user.email"), string(user.Email)).
		Relation("User").
		Relation("ListItems").
		Relation("ListItems.Product").
		Relation("ListItems.Category").
		Relation("ListItems.Product.Categories").
		Relation("ListItems.Product.Categories.Category").
		Scan(ctx)
	if err != nil {
		return []domain.List{}, fmt.Errorf("repositories.ListRepository.GetAll: failed to get lists for user: %v:\n%w", user, err)
	}

	lists, err := mappers.Map(dbLists, mappers.DbListToDomain)
	if err != nil {
		return []domain.List{}, fmt.Errorf("repositories.ListRepository.GetAll: failed to map lists\n%w", err)
	}

	return lists, nil
}

func (r *ListRepository) Create(ctx context.Context, list domain.List) error {
	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbList := postgresql.List{
			Id:        list.Id,
			Name:      string(list.Name),
			UserId:    list.User.Id,
			ListItems: []postgresql.ListItem{},
		}

		_, err := tx.NewInsert().
			Model(&dbList).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("repositories.ListRepository.Create: failed to create list: %v:\n%w", list, err)
		}
		return nil
	})
}

func (r *ListRepository) Update(ctx context.Context, list domain.List) error {
	existing, err := r.Get(ctx, list.User, list.Name)
	if err != nil {
		return fmt.Errorf("repositories.ListRepository.Update: failed to get list: %v\n%w", list, err)
	}

	if existing.Id == uuid.Nil {
		return fmt.Errorf("repositories.ListRepository.Update:\n%w", ErrListNotFound)
	}

	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbList := postgresql.List{
			Id:        list.Id,
			Name:      string(list.Name),
			UserId:    list.User.Id,
			ListItems: []postgresql.ListItem{},
		}

		_, err := tx.NewUpdate().
			Model(&dbList).
			WherePK().
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("repositories.ListRepository.Update: failed to update list: %v:\n%w", list, err)
		}
		return nil
	})
}
