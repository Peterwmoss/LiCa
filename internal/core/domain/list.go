package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidListName = errors.New("invalid list name")
)

type List struct {
	Id               uuid.UUID
	Name             ListName
	Items            []ListItem
	CategoryOrdering map[int]Category
	User             User
}

func CreateList(name ListName, user User) List {
	return List{
		Id:               uuid.New(),
		Name:             name,
		User:             user,
		Items:            []ListItem{},
		CategoryOrdering: make(map[int]Category),
	}
}

func NewList(id uuid.UUID, name ListName, items []ListItem, categoryOrdering map[int]Category, user User) List {
	return List{
		Id:               id,
		Name:             name,
		Items:            items,
		CategoryOrdering: categoryOrdering,
		User:             user,
	}
}

type ListName string

func NewListName(name string) (ListName, error) {
	if name == "" {
		return "", ErrInvalidListName
	}

	return ListName(name), nil
}
