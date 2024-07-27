package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql/mappers"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) ports.UserRepository {
	return &UserRepository{
		db,
	}
}

func (u *UserRepository) Create(ctx context.Context, user domain.User) error {
	return u.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		slog.Debug(fmt.Sprintf("creating user with email: %s", string(user.Email)))

		dbUser := postgresql.User{
      Id: user.Id,
      Email: string(user.Email),
    }
		_, err := tx.NewInsert().
			Model(&dbUser).
			Exec(ctx)

    return err
	})
}

func (u *UserRepository) GetByEmail(ctx context.Context, email domain.Email) (domain.User, error) {
	slog.Debug(fmt.Sprintf("getting user with email: %s", string(email)))

	dbUser := postgresql.User{}
	err := u.db.NewSelect().
		Model(&dbUser).
		Where("? like ?", bun.Ident("u.email"), string(email)).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			slog.Debug(fmt.Sprintf("user with email: %s not found, returning empty user", email))
			return domain.User{}, nil
		}

		return domain.User{}, err
	}

	return mappers.DbUserToDomain(dbUser)
}

func (u *UserRepository) UpdateEmail(ctx context.Context, id uuid.UUID, email domain.Email) error {
	panic("unimplemented")
}
