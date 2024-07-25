package mappers

import (
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

func DbUserToDomain(dbUser postgresql.User) (domain.User, error) {
	domainEmail, err := domain.NewEmail(dbUser.Email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Id:    dbUser.Id,
		Email: domainEmail,
	}, nil
}
