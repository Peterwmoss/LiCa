package mappers

import (
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

func DbCategoryToDomain(dbCategory postgresql.Category) (domain.Category, error) {
	domainName, err := domain.NewCategoryName(dbCategory.Name)
	if err != nil {
		return domain.Category{}, err
	}

	domainUser, err := DbUserToDomain(dbCategory.User)
	if err != nil {
		return domain.Category{}, err
	}

	return domain.Category{
		Id:   dbCategory.Id,
		Name: domainName,
		User: domainUser,
	}, nil
}
