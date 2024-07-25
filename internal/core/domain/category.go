package domain

import (
	"errors"

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
		return "", ErrInvalidCategoryName
	}

	return CategoryName(name), nil
}
