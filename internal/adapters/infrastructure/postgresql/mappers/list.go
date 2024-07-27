package mappers

import (
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

func DbListToDomain(dbList postgresql.List) (domain.List, error) {
	domainUser, err := DbUserToDomain(dbList.User)
	if err != nil {
		return domain.List{}, fmt.Errorf("DbListToDomain: Failed to map user: %w", err)
	}

	domainName, err := domain.NewListName(dbList.Name)
	if err != nil {
		return domain.List{}, fmt.Errorf("DbListToDomain: Failed to create list name: %w", err)
	}

	domainItems := make([]domain.ListItem, len(dbList.ListItems))
	for i, item := range dbList.ListItems {
		domainItem, err := DbListItemToDomain(item)
		if err != nil {
			return domain.List{}, fmt.Errorf("DbListToDomain: Failed to map list item: %v, %w", item, err)
		}

		domainItems[i] = domainItem
	}

	categoryOrderings := make(map[int]domain.Category)
	return domain.NewList(dbList.Id, domainName, domainItems, categoryOrderings, domainUser), nil
}
