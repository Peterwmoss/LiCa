package mappers

import (
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

func DbUserToDomain(dbUser postgresql.User) (domain.User, error) {
	domainEmail, err := domain.NewEmail(dbUser.Email)
	if err != nil {
    return domain.User{}, fmt.Errorf("mappers.DbUserToDomain: Failed to create email: %v:. Error: %w", dbUser.Email, err)
	}

	return domain.User{
		Id:    dbUser.Id,
		Email: domainEmail,
	}, nil
}
