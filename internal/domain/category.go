package domain

import (
	"context"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/uptrace/bun"
)

type (
	Category struct {
		Id       int
		Name     string
		Order    int
		IsCustom bool
	}

	CategoryService interface {
		GetAll(User) ([]Category, error)
		ToDomain(database.Category) Category
		FromDomain(Category, User) database.Category

		GetById(id int) (*Category, error)
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
		Where("user_id = ? or user_id IS NULL", user.id).
		Relation("Orders").
		Scan(service.ctx)
	if err != nil {
		return nil, fmt.Errorf("GetAll: failed to get categories for user: %d: %w", user.id, err)
	}

	categories := make([]Category, len(arr))
	for i := range arr {
		categories[i] = service.ToDomain(arr[i])
	}

	return categories, nil
}

func (service categoryService) GetById(id int) (*Category, error) {
	category := database.Category{}

	err := service.db.NewSelect().
		Model(&category).
		Where("id = ?", id).
		Relation("Orders").
		Limit(1).
		Scan(service.ctx)
	if err != nil {
		return nil, fmt.Errorf("GetById: failed to get category with id: %d: %w", id, err)
	}

	domainCategory := service.ToDomain(category)

	return &domainCategory, nil
}

func (categoryService) ToDomain(category database.Category) Category {
	order := 0
	if len(category.Orders) > 0 {
		order = category.Orders[0].Order
	}
	return Category{
		Id:       category.Id,
		Name:     category.Name,
		Order:    order,
		IsCustom: category.IsCustom,
	}
}

func (categoryService) FromDomain(category Category, user User) database.Category {
	return database.Category{
		Id:   category.Id,
		Name: category.Name,
		Orders: []database.CategoryOrder{
			{
				CategoryId: category.Id,
				UserId:     user.id,
				Order:      category.Order,
			},
		},
		UserId:   &user.id,
		IsCustom: category.IsCustom,
	}
}
