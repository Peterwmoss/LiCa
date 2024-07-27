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
	ErrProductNotFound = errors.New("product does not exist")
)

type ProductRepository struct {
	db *bun.DB
}

func NewProductRepository(db *bun.DB) ports.ProductRepository {
	return &ProductRepository{
		db,
	}
}

func (r *ProductRepository) GetById(ctx context.Context, user domain.User, id uuid.UUID) (domain.Product, error) {
	dbProduct := postgresql.Product{}

	err := r.db.NewSelect().
		Model(&dbProduct).
		Where("id = ?", id).
		Where("? like ?", bun.Ident("user.email"), string(user.Email)).
		WhereOr("? IS NULL", bun.Ident("user_id")).
		Relation("User").
		Relation("Categories").
		Relation("Categories.Category").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return domain.Product{}, err
	}

	return mappers.DbProductToDomain(dbProduct)
}

func (r *ProductRepository) Get(ctx context.Context, user domain.User, name domain.ProductName) (domain.Product, error) {
	dbProduct := postgresql.Product{}

	err := r.db.NewSelect().
		Model(&dbProduct).
		Where("name = ?", name).
		Where("? like ?", bun.Ident("user.email"), string(user.Email)).
		WhereOr("? IS NULL", bun.Ident("user_id")).
		Relation("User").
		Relation("Categories").
		Relation("Categories.Category").
		Limit(1).
		Scan(ctx)
	if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return domain.Product{}, nil
    }
		return domain.Product{}, err
	}

	return mappers.DbProductToDomain(dbProduct)
}

func (r *ProductRepository) GetAll(ctx context.Context, user domain.User) ([]domain.Product, error) {
	var dbProducts []postgresql.Product

	err := r.db.NewSelect().
		Model(&dbProducts).
		Where("? like ?", bun.Ident("user.email"), string(user.Email)).
		WhereOr("? IS NULL", bun.Ident("user_id")).
		Relation("User").
		Relation("Categories").
		Relation("Categories.Category").
		Scan(ctx)
	if err != nil {
		return []domain.Product{}, err
	}

	return mappers.Map(dbProducts, mappers.DbProductToDomain)
}

func (r *ProductRepository) Create(ctx context.Context, product domain.Product) error {
	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbProduct := postgresql.Product{
			Id:         product.Id,
			Name:       string(product.Name),
			UserId:     product.User.Id,
			Categories: []postgresql.ProductCategories{},
		}

		_, err := tx.NewInsert().
			Model(&dbProduct).
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *ProductRepository) Update(ctx context.Context, product domain.Product) error {
	existing, err := r.Get(ctx, product.User, product.Name)
	if err != nil {
		return err
	}

	if existing.Id == uuid.Nil {
		return ErrProductNotFound
	}

	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbProduct := postgresql.Product{
			Id:     product.Id,
			Name:   string(product.Name),
			UserId: product.User.Id,
		}

		_, err := tx.NewUpdate().
			Model(&dbProduct).
			WherePK().
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
