package mappers

import (
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

func DbListItemToDomain(dbItem postgresql.ListItem) (domain.ListItem, error) {
	domainProduct, err := DbProductToDomain(dbItem.Product)
	if err != nil {
		return domain.ListItem{}, err
	}

	domainAmount, err := domain.NewAmount(dbItem.Amount)
	if err != nil {
		return domain.ListItem{}, err
	}

	domainUnit, err := domain.NewUnit(dbItem.Unit)
	if err != nil {
		return domain.ListItem{}, err
	}

	domainCategory, err := DbCategoryToDomain(dbItem.Category)
	if err != nil {
		return domain.ListItem{}, err
	}

	return domain.NewListItem(dbItem.Id, domainProduct, domainAmount, domainUnit, domainCategory), nil
}
