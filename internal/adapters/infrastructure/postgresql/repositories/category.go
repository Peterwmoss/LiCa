package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql/mappers"
	"github.com/Peterwmoss/LiCa/internal/core"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var (
	ErrCategoryNotFound = errors.New("category does not exist")
)

type CategoryRepository struct {
	db *bun.DB
}

func NewCategoryRepository(db *bun.DB) ports.CategoryRepository {
	return &CategoryRepository{
		db,
	}
}

func (r *CategoryRepository) GetById(ctx context.Context, user domain.User, id uuid.UUID) (domain.Category, error) {
	dbCategory := postgresql.Category{}

	err := r.db.NewSelect().
		Model(&dbCategory).
		Where("id = ?", id).
		Where("(? like ? OR ? IS NULL)", bun.Ident("user.email"), string(user.Email), bun.Ident("user_id")).
		Relation("User").
		Relation("Categories").
		Relation("Categories.Category").
		Limit(1).
		Scan(ctx)
	if err != nil {
		err = fmt.Errorf("repositories.CategoryRepository.GetById: failed to get category with id: %s:. Error: %w", id, err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Category{}, fmt.Errorf("%w. Error: %w", core.ErrNotFound, err)
		}

		return domain.Category{}, err
	}

	category, err := mappers.DbCategoryToDomain(dbCategory)
	if err != nil {
		return domain.Category{}, fmt.Errorf("repositories.CategoryRepository.GetById: failed to map category. Error: %w", err)
	}

	return category, nil
}

func (r *CategoryRepository) Get(ctx context.Context, user domain.User, name domain.CategoryName) (domain.Category, error) {
	dbCategory := postgresql.Category{}

	err := r.db.NewSelect().
		Model(&dbCategory).
		Where("name = ?", name).
		Where("(? like ? OR ? IS NULL)", bun.Ident("user.email"), string(user.Email), bun.Ident("user_id")).
		Relation("User").
		Limit(1).
		Scan(ctx)
	if err != nil {
		err = fmt.Errorf("repositories.CategoryRepository.Get: failed to get category with name: %s:. Error: %w", name, err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Category{}, fmt.Errorf("%w. Error: %w", core.ErrNotFound, err)
		}

		return domain.Category{}, err
	}

	category, err := mappers.DbCategoryToDomain(dbCategory)
	if err != nil {
		return domain.Category{}, fmt.Errorf("repositories.CategoryRepository.Get: failed to map category. Error: %w", err)
	}

	return category, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context, user domain.User) ([]domain.Category, error) {
	var dbCategories []postgresql.Category

	err := r.db.NewSelect().
		Model(&dbCategories).
		Where("(? like ? OR ? IS NULL)", bun.Ident("user.email"), string(user.Email), bun.Ident("user_id")).
		Relation("User").
		Scan(ctx)
	if err != nil {
		return []domain.Category{}, fmt.Errorf("repositories.CategoryRepository.GetAll: failed to get categories for user: %v:. Error: %w", user, err)
	}

	slog.Debug("Found categories", "categories", dbCategories)
	categories, err := mappers.Map(dbCategories, mappers.DbCategoryToDomain)
	if err != nil {
		return []domain.Category{}, fmt.Errorf("repositories.CategoryRepository.GetAll: failed to map categories. Error: %w", err)
	}

	return categories, nil
}

func (r *CategoryRepository) Create(ctx context.Context, category domain.Category) error {
	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbCategory := postgresql.Category{
			Id:     category.Id,
			Name:   string(category.Name),
			UserId: category.User.Id,
		}

		_, err := tx.NewInsert().
			Model(&dbCategory).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("repositories.CategoryRepository.Create: Failed to create category: %v:. Error: %w", category, err)
		}
		return nil
	})
}

func (r *CategoryRepository) Update(ctx context.Context, category domain.Category) error {
	existing, err := r.Get(ctx, category.User, category.Name)
	if err != nil {
		return fmt.Errorf("repositories.CategoryRepository.Update: Failed to get category: %v:. Error: %w", category, err)
	}

	if existing.Id == uuid.Nil {
		return fmt.Errorf("repositories.CategoryRepository.Update:. Error: %w", ErrCategoryNotFound)
	}

	return r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		dbCategory := postgresql.Category{
			Id:     category.Id,
			Name:   string(category.Name),
			UserId: category.User.Id,
		}

		_, err := tx.NewUpdate().
			Model(&dbCategory).
			WherePK().
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("repositories.CategoryRepository.Update: Failed to update category: %v:. Error: %w", category, err)
		}
		return nil
	})
}
