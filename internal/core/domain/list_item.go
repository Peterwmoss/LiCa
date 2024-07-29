package domain

import (
	"errors"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/core"
	"github.com/google/uuid"
)

var (
	ErrInvalidListItemAmount = errors.New("invalid amount")
)

type ListItem struct {
	Id       uuid.UUID
	Product  Product
	Amount   Amount
	Unit     Unit
	Category Category
	List     List
}

func NewListItem(id uuid.UUID, list List, product Product, amount Amount, unit Unit, category Category) ListItem {
	return ListItem{
		Id:       id,
		List:     list,
		Product:  product,
		Amount:   amount,
		Unit:     unit,
		Category: category,
	}
}

func CreateListItem(list List, product Product, amount Amount, unit Unit, category Category) ListItem {
	return ListItem{
		Id:       uuid.New(),
		List:     list,
		Product:  product,
		Amount:   amount,
		Unit:     unit,
		Category: category,
	}
}

type Amount float32

func NewAmount(amount float32) (Amount, error) {
	if amount < 0 {
		return 0, fmt.Errorf("domain.NewAmount: amount must be positive\n%w\n%w", ErrInvalidListItemAmount, core.ErrValidation)
	}

	return Amount(amount), nil
}

type Unit string

func NewUnit(unit string) (Unit, error) {
	return Unit(unit), nil
}
