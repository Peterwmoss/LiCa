package domain

import (
	"errors"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/core"
	"github.com/google/uuid"
)

var (
	ErrInvalidCategoryName = errors.New("invalid category name")
)

type Category struct {
	Id   uuid.UUID
	Name CategoryName
	User User
}

func (c Category) String() string {
  return fmt.Sprintf("Category{Name:%v, User:%v}", c.Name, c.User)
}

func CreateCategory(name CategoryName, user User) Category {
	return Category{
		Id:   uuid.New(),
		Name: name,
		User: user,
	}
}

func NewCategory(id uuid.UUID, name CategoryName, user User) Category {
	return Category{
		Id:   id,
		Name: name,
		User: user,
	}
}

type CategoryName string

func NewCategoryName(name string) (CategoryName, error) {
	if name == "" {
		return "", fmt.Errorf("domain.NewCategoryName: name must not be empty. Error: %w. Error: %w", ErrInvalidCategoryName, core.ErrValidation)
	}

	return CategoryName(name), nil
}
