package services

import (
	"context"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
)

type ListItemService struct {
	productService  ports.ProductService
	categoryService ports.CategoryService
	listService     ports.ListService
	repository      ports.ListItemRepository
}

func NewListItemService(repository ports.ListItemRepository, ps ports.ProductService, cs ports.CategoryService, ls ports.ListService) ports.ListItemService {
	return &ListItemService{
		productService:  ps,
		categoryService: cs,
		listService:     ls,
		repository:      repository,
	}
}

// Add implements ports.ListItemService.
func (l *ListItemService) Add(ctx context.Context, createItem ports.ListItemCreate) (domain.ListItem, error) {
	list, err := l.listService.Get(ctx, createItem.ListName)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("services.ListItemService.Create: Error: %w", err)
	}

	product, err := l.productService.Get(ctx, createItem.ProductName)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("services.ListItemService.Create: Error: %w", err)
	}

	if product.Id == uuid.Nil {
		product, err = l.productService.Create(ctx, createItem.ProductName)
		if err != nil {
			return domain.ListItem{}, fmt.Errorf("services.ListItemService.Create: Error: %w", err)
		}
	}

	amount, err := domain.NewAmount(createItem.Amount)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("services.ListItemService.Create: Error: %w", err)
	}

	unit, err := domain.NewUnit("stk")
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("services.ListItemService.Create: Error: %w", err)
	}

	category, err := l.categoryService.Get(ctx, createItem.CategoryName)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("services.ListItemService.Create: Error: %w", err)
	}

	domainItem := domain.CreateListItem(list, product, amount, unit, category)

	user := ctx.Value("user").(domain.User)

	created, err := l.repository.Create(ctx, user, domainItem)
	if err != nil {
		return domain.ListItem{}, fmt.Errorf("services.ListItemService.Create: Error: %w", err)
	}

	return created, nil
}

// Remove implements ports.ListItemService.
func (l *ListItemService) Remove(ctx context.Context, id uuid.UUID) error {
	user := ctx.Value("user").(domain.User)

	err := l.repository.Remove(ctx, user, id)
	if err != nil {
		return fmt.Errorf("services.ListItemService.Remove: Error: %w", err)
	}

	return nil
}

// Update implements ports.ListItemService.
func (l *ListItemService) Update(ctx context.Context, updateItem ports.ListItemUpdate) error {
	user := ctx.Value("user").(domain.User)

	existing, err := l.repository.GetById(ctx, user, updateItem.Id)
	if err != nil {
		return fmt.Errorf("services.ListItemService.Update: Error: %w", err)
	}

	amount, err := domain.NewAmount(updateItem.Amount)
	if err != nil {
		return fmt.Errorf("services.ListItemService.Update: Error: %w", err)
	}

	category, err := l.categoryService.Get(ctx, updateItem.CategoryName)
	if err != nil {
		return fmt.Errorf("services.ListItemService.Update: Error: %w", err)
	}

	domainItem := domain.NewListItem(updateItem.Id, existing.List, existing.Product, amount, existing.Unit, category)

	_, err = l.repository.Update(ctx, user, domainItem)
	if err != nil {
		return fmt.Errorf("services.ListItemService.Update: Error: %w", err)
	}

	return nil
}

// GetAll implements ports.ListItemService.
func (l *ListItemService) GetAll(ctx context.Context, listName string) ([]domain.ListItem, error) {
	domainListName, err := domain.NewListName(listName)
	if err != nil {
		return []domain.ListItem{}, fmt.Errorf("services.ListItemService.GetAll: Error: %w", err)
	}

	user := ctx.Value("user").(domain.User)

	items, err := l.repository.GetAll(ctx, user, domainListName)
	if err != nil {
		return []domain.ListItem{}, fmt.Errorf("services.ListItemService.GetAll: Error: %w", err)
	}

	return items, nil
}
