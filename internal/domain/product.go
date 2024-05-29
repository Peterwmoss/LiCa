package domain

import (
	"context"
	"errors"

	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/jackc/pgerrcode"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type (
	Product struct {
		Id       int
		Name     string
		Category Category
		IsCustom bool
	}

	ProductService interface {
		ToDomain(database.Product) Product
		ToDatabase(Product) database.Product
		CreateIfNotExists(name string, category *Category, user User) (*Product, error)
		Get(name string, category *Category, user User) (*Product, error)
	}

	productService struct {
		db              *bun.DB
		ctx             context.Context
		categoryService CategoryService
	}
)

func NewProductService(db *bun.DB, ctx context.Context, categoryService CategoryService) ProductService {
	return &productService{db, ctx, categoryService}
}

func (svc productService) CreateIfNotExists(name string, category *Category, user User) (*Product, error) {
	if name == "" {
		return nil, EmptyNameError
	}

	existing, err := svc.Get(name, category, user)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return existing, nil
	}

	product := database.Product{
		Name:       name,
		CategoryId: category.Id,
		UserId:     &user.id,
	}

	_, err = svc.db.NewInsert().
		Model(&product).
		Exec(svc.ctx)
	if err != nil {
		errStatusCode := err.(pgdriver.Error).Field('C')
		if errStatusCode == pgerrcode.UniqueViolation {
			return nil, errors.Join(err, UniqueViolationError)
		}

		return nil, err
	}

	return svc.Get(name, category, user)
}

func (svc productService) Get(name string, category *Category, user User) (*Product, error) {
	databaseProduct := database.Product{}

	err := svc.db.NewSelect().
		Model(&databaseProduct).
		Where("name = ?", name).
		Where("category_id = ?", category.Id).
		Where("user_id = ? OR user_id IS NULL", user.id).
		Relation("Category").
		Relation("Category.Orders").
		Limit(1).
		Scan(svc.ctx)
	if err != nil {
		return nil, err
	}

  if databaseProduct.Id == 0 {
    log.Debug().Msgf("no product found for name: \"%s\", categoryId: \"%d\" and userId: \"%d\" ", name, category.Id, user.id)
		return nil, nil
  }

	product := svc.ToDomain(databaseProduct)

	log.Debug().Msgf("found product: %v", product)

	return &product, nil
}

func (svc productService) ToDatabase(product Product) database.Product {
	// TODO: implement
	return database.Product{}
}

func (svc productService) ToDomain(product database.Product) Product {
	category := svc.categoryService.ToDomain(*product.Category)
	return Product{
		Id:       product.Id,
		Name:     product.Name,
		Category: category,
		IsCustom: product.IsCustom,
	}
}
