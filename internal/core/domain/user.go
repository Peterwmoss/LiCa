package domain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Peterwmoss/LiCa/internal/core"
	"github.com/google/uuid"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type User struct {
	Id    uuid.UUID
	Email Email
}

func CreateUser(email Email) User {
	return User{
		Id:    uuid.New(),
		Email: email,
	}
}

func NewUser(id uuid.UUID, email Email) User {
	return User{
		Id:    id,
		Email: email,
	}
}

type Email string

func NewEmail(email string) (Email, error) {
	if !strings.Contains(email, "@") {
		return "", fmt.Errorf("domain.NewEmail: email must contain '@': '%s'\n%w\n%w", email, ErrInvalidEmail, core.ErrValidation)
	}

	return Email(email), nil
}
