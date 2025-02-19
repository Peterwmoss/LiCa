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

func (l ListItem) String() string {
  categoryId := l.Category.Id.String()
  if categoryId == uuid.Nil.String() {
    categoryId = "(nil)"
  }

  productId := l.Product.Id.String()
  if productId == uuid.Nil.String() {
    productId = "(nil)"
  }

  listId := l.List.Id.String()
  if listId == uuid.Nil.String() {
    listId = "(nil)"
  }

  return fmt.Sprintf("ListItem{Id:%v, List:%v, Category:%v, Product:%v, Amount:%f, Unit:%v}", l.Id, listId, categoryId, productId, l.Amount, l.Unit)
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
		return 0, fmt.Errorf("domain.NewAmount: amount must be positive. Error: %w. Error: %w", ErrInvalidListItemAmount, core.ErrValidation)
	}

	return Amount(amount), nil
}

type Unit string

func NewUnit(unit string) (Unit, error) {
	return Unit(unit), nil
}
