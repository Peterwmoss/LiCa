package mappers

import (
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/google/uuid"
)

func DbListItemToDomain(dbItem postgresql.ListItem) (domain.ListItem, error) {
	domainProduct, err := DbProductToDomain(dbItem.Product)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("mappers.DbListItemToDomain: Failed to map product: %v:. Error: %w", dbItem.Product, err)
	}

	domainAmount, err := domain.NewAmount(dbItem.Amount)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("mappers.DbListItemToDomain: Failed to create amount: %v:. Error: %w", dbItem.Amount, err)
	}

	domainUnit, err := domain.NewUnit(dbItem.Unit)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("mappers.DbListItemToDomain: Failed to create unit: %v:. Error: %w", dbItem.Unit, err)
	}

	domainCategory, err := DbCategoryToDomain(dbItem.Category)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("mappers.DbListItemToDomain: Failed to map category: %v:. Error: %w", dbItem.Category, err)
	}

	var domainList domain.List
	if dbItem.List.Id != uuid.Nil {
		domainList, err = DbListToDomain(dbItem.List)
		if err != nil {
			return domain.ListItem{}, fmt.Errorf("mappers.DbListItemToDomain: Failed to map list: %v:. Error: %w", dbItem.List, err)
		}
	}

	return domain.NewListItem(dbItem.Id, domainList, domainProduct, domainAmount, domainUnit, domainCategory), nil
}
