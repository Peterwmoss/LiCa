package domain

import (
	"errors"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/core"
	"github.com/google/uuid"
)

var (
	ErrInvalidProductName = errors.New("invalid product name")
)

type Product struct {
	Id         uuid.UUID
	Name       ProductName
	Categories []Category
	IsCustom   bool
	User       User
}

func CreateProduct(name ProductName, categories []Category, user User) Product {
	return Product{
		Id:         uuid.New(),
		Name:       name,
		Categories: categories,
		IsCustom:   user.Id != uuid.Nil,
		User:       user,
	}
}

func NewProduct(id uuid.UUID, name ProductName, categories []Category, isCustom bool, user User) Product {
	return Product{
		Id:         id,
		Name:       name,
		Categories: categories,
		IsCustom:   isCustom,
		User:       user,
	}
}

type ProductName string

func NewProductName(name string) (ProductName, error) {
	if name == "" {
		return "", fmt.Errorf("domain.NewProductName: name must not be empty\n%w\n%w", ErrInvalidProductName, core.ErrValidation)
	}

	return ProductName(name), nil
}
