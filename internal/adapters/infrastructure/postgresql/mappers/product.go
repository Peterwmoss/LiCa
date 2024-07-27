package mappers

import (
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/google/uuid"
)

func DbProductToDomain(dbProduct postgresql.Product) (domain.Product, error) {
	// Ignore error in case no user is connected to product
	domainUser, _ := DbUserToDomain(dbProduct.User)

	domainCategories := make([]domain.Category, len(dbProduct.Categories))
	for i, category := range dbProduct.Categories {
		domainCategory, err := DbCategoryToDomain(category.Category)
		if err != nil {
			return domain.Product{}, err
		}
		domainCategories[i] = domainCategory
	}

	domainName, err := domain.NewProductName(dbProduct.Name)
	if err != nil {
		return domain.Product{}, err
	}

	isCustom := domainUser.Id != uuid.Nil

	return domain.NewProduct(dbProduct.Id, domainName, domainCategories, isCustom, domainUser), nil
}
