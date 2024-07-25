package repositories

import (
	"context"
	"database/sql"
	"errors"

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

func (repo *ListRepository) Get(ctx context.Context, email domain.Email, name domain.ListName) (domain.List, error) {
	dbList := postgresql.List{}

	err := repo.db.NewSelect().
		Model(&dbList).
		Where("name = ?", name).
		Where("u.email = ?", email).
		Relation("User").
		Relation("ListItems").
		Relation("ListItems.Product").
		Relation("ListItems.Category").
		Relation("ListItems.Product.Categories").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return domain.List{}, err
	}

	return mappers.DbListToDomain(dbList)
}

func (l *ListRepository) GetAllByEmail(ctx context.Context, email domain.Email) ([]domain.List, error) {
	panic("unimplemented")
}

func (l *ListRepository) Create(ctx context.Context, list domain.List) error {
	return l.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
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
			return err
		}

		return nil
	})
}

func (l *ListRepository) Update(ctx context.Context, list domain.List) error {
	existing, err := l.Get(ctx, list.User.Email, list.Name)
	if err != nil {
		return err
	}

	if existing.Id == uuid.Nil {
		return ErrListNotFound
	}

	return l.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
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
			return err
		}

		return nil
	})
}
