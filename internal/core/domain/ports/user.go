package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/google/uuid"
)

type UserService interface {
	// Create a new user.
	//
	// Returns the created user and no error if user doesn't exist.
	// If the user exists nil is returned along with an error.
	Create(ctx context.Context, email string) (domain.User, error)

	// Get a user by email.
	//
	// Returns nil and no error if no user with the provided email is not found
	Get(ctx context.Context, email string) (domain.User, error)
}

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetById(ctx context.Context, id uuid.UUID) (domain.User, error)
	UpdateEmail(ctx context.Context, id uuid.UUID, email domain.Email) error
	GetByEmail(ctx context.Context, email domain.Email) (domain.User, error)
}
