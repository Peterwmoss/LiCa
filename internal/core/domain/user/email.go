package user

import (
	"errors"
	"strings"

	"github.com/Peterwmoss/LiCa/internal"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Email string

func NewEmail(email string) (Email, error) {
	if !strings.Contains(email, "@") {
		return "", errors.Join(core.ErrBadRequest, ErrInvalidEmail)
	}

	return Email(email), nil
}
