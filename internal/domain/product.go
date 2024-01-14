package domain

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/uptrace/bun"
)

type (
	Product struct {
		id       int64
		Name     string
		Category Category
		IsCustom bool
	}

	ProductService interface {
		ToDomain(database.Product) Product
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

func (svc productService) ToDomain(product database.Product) Product {
	category := svc.categoryService.ToDomain(*product.Category)
	return Product{
		id:       product.Id,
		Name:     product.Name,
		Category: category,
		IsCustom: product.IsCustom,
	}
}
