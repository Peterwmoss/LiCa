package services

import (
	"context"
	"errors"

	"github.com/Peterwmoss/LiCa/internal/core"
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
		return domain.ListItem{}, err
	}

	product, err := l.productService.Get(ctx, createItem.ProductName)
	if err != nil {
		return domain.ListItem{}, err
	}

	if product.Id == uuid.Nil {
		product, err = l.productService.Create(ctx, createItem.ProductName)
		if err != nil {
			return domain.ListItem{}, err
		}
	}

	amount, err := domain.NewAmount(createItem.Amount)
	if err != nil {
		return domain.ListItem{}, errors.Join(err, core.ErrValidation)
	}

	unit, err := domain.NewUnit("stk")
	if err != nil {
		return domain.ListItem{}, errors.Join(err, core.ErrValidation)
	}

	category, err := l.categoryService.Get(ctx, createItem.CategoryName)
	if err != nil {
		return domain.ListItem{}, err
	}

	if category.Id == uuid.Nil {
		category, err = l.categoryService.Create(ctx, createItem.CategoryName)
		if err != nil {
			return domain.ListItem{}, err
		}
	}

	domainItem := domain.CreateListItem(list, product, amount, unit, category)

	user := ctx.Value("user").(domain.User)

	return l.repository.Create(ctx, user, domainItem)
}

// Remove implements ports.ListItemService.
func (l *ListItemService) Remove(ctx context.Context, id uuid.UUID) error {
	user := ctx.Value("user").(domain.User)
	return l.repository.Remove(ctx, user, id)
}

// Update implements ports.ListItemService.
func (l *ListItemService) Update(ctx context.Context, updateItem ports.ListItemUpdate) error {
	user := ctx.Value("user").(domain.User)

	existing, err := l.repository.GetById(ctx, user, updateItem.Id)
	if err != nil {
		return err
	}

	product, err := l.productService.GetById(ctx, existing.Product.Id)
	if err != nil {
		return err
	}

	amount, err := domain.NewAmount(updateItem.Amount)
	if err != nil {
		return errors.Join(err, core.ErrValidation)
	}

	category, err := l.categoryService.Get(ctx, updateItem.CategoryName)
	if err != nil {
		return err
	}

	if category.Id == uuid.Nil {
		category, err = l.categoryService.Create(ctx, updateItem.CategoryName)
		if err != nil {
			return err
		}
	}

	domainItem := domain.NewListItem(updateItem.Id, existing.List, product, amount, existing.Unit, category)

	_, err = l.repository.Update(ctx, user, domainItem)
	return err
}

// GetAll implements ports.ListItemService.
func (l *ListItemService) GetAll(ctx context.Context, listName string) ([]domain.ListItem, error) {
	domainListName, err := domain.NewListName(listName)
	if err != nil {
		return []domain.ListItem{}, errors.Join(err, core.ErrValidation)
	}

	user := ctx.Value("user").(domain.User)

	return l.repository.GetAll(ctx, user, domainListName)
}
