package domain

import (
	"errors"

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
}

func NewListItem(id uuid.UUID, product Product, amount Amount, unit Unit, category Category) ListItem {
	return ListItem{
		Id:       id,
		Product:  product,
		Amount:   amount,
		Unit:     unit,
		Category: category,
	}
}

func CreateListItem(product Product, amount Amount, unit Unit, category Category) ListItem {
	return ListItem{
		Id:       uuid.New(),
		Product:  product,
		Amount:   amount,
		Unit:     unit,
		Category: category,
	}
}

type Amount float32

func NewAmount(amount float32) (Amount, error) {
	if amount < 1 {
		return 0, ErrInvalidListItemAmount
	}

	return Amount(amount), nil
}

type Unit string

func NewUnit(unit string) (Unit, error) {
	return Unit(unit), nil
}
