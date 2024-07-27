package ports

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/google/uuid"
)

type UserService interface {
	// Create a new user.
	Create(ctx context.Context, email string) (domain.User, error)

	// Get a user by email.
	//
	// Doesn't error if email doesn't exist
	Get(ctx context.Context, email string) (domain.User, error)
}

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	UpdateEmail(ctx context.Context, id uuid.UUID, email domain.Email) error
	GetByEmail(ctx context.Context, email domain.Email) (domain.User, error)
}
