package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain/user"
	"github.com/google/uuid"
)

type UserPort interface {
	// Create a new user.
	//
	// Returns the created user and no error if user doesn't exist.
	// If the user exists nil is returned along with an error.
	Create(ctx context.Context, email string) (*user.User, error)

	// Get a user by email.
	//
	// Returns nil and no error if no user with the provided email is not found
	Get(ctx context.Context, email string) (*user.User, error)
}

type UserCommandRepository interface {
	Create(ctx context.Context, user *user.User) error
	Get(ctx context.Context, id uuid.UUID) (*user.User, error)
	UpdateEmail(ctx context.Context, id uuid.UUID, email user.Email) error
}

type UserQueryRepository interface {
	UserByEmail(ctx context.Context, email user.Email) (*user.User, error)
}
