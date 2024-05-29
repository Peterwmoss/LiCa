package category

import "github.com/google/uuid"

type Category struct {
	Id       uuid.UUID
	IsCustom bool
	Name     Name
}

func New(id uuid.UUID, isCustom bool, name string) (*Category, error) {
	domainName, err := NewName(name)
	if err != nil {
		return nil, err
	}

	return &Category{
		Id:       id,
		IsCustom: isCustom,
		Name:     domainName,
	}, nil
}

func Create(name Name) *Category {
	return &Category{
		Id:       uuid.New(),
		IsCustom: true,
		Name:     name,
	}
}
