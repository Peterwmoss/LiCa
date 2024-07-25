package mappers

import (
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

func DbProductToDomain(dbProduct postgresql.Product) (domain.Product, error) {
	domainUser, err := DbUserToDomain(dbProduct.User)
	if err != nil {
		return domain.Product{}, err
	}

	domainCategories := make([]domain.Category, len(dbProduct.Categories))
	for i, category := range dbProduct.Categories {
		domainCategory, err := DbCategoryToDomain(category)
		if err != nil {
			return domain.Product{}, err
		}
		domainCategories[i] = domainCategory
	}

	domainName, err := domain.NewProductName(dbProduct.Name)
	if err != nil {
		return domain.Product{}, err
	}

	return domain.NewProduct(dbProduct.Id, domainName, domainCategories, dbProduct.IsCustom, domainUser), nil
}
