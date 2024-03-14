package domain

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/uptrace/bun"
)

type (
	Category struct {
		id       int
		Name     string
		Order    int
		IsCustom bool
	}

	CategoryService interface {
		GetAll(User) ([]Category, error)
		ToDomain(database.Category) Category
		FromDomain(Category, User) database.Category
	}

	categoryService struct {
		db  *bun.DB
		ctx context.Context
	}
)

func NewCategoryService(db *bun.DB, ctx context.Context) CategoryService {
	return &categoryService{db, ctx}
}

func (service categoryService) GetAll(user User) ([]Category, error) {
	arr := []database.Category{}

	err := service.db.NewSelect().
		Model(&arr).
		Where("user_id = ?", user.id).
		Relation("Orders").
		Scan(service.ctx)
	if err != nil {
		return nil, err
	}

	categories := make([]Category, len(arr))
	for i := range arr {
		categories[i] = service.ToDomain(arr[i])
	}

	return categories, nil
}

func (categoryService) ToDomain(category database.Category) Category {
	order := 0
	if len(category.Orders) > 0 {
		order = category.Orders[0].Order
	}
	return Category{
		id:       category.Id,
		Name:     category.Name,
		Order:    order,
		IsCustom: category.IsCustom,
	}
}

func (categoryService) FromDomain(category Category, user User) database.Category {
	return database.Category{
		Id:   category.id,
		Name: category.Name,
		Orders: []database.CategoryOrder{
			{
				CategoryId: category.id,
				UserId:     user.id,
				Order:      category.Order,
			},
		},
		UserId:   &user.id,
		IsCustom: category.IsCustom,
	}
}
