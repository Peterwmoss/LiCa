package services

import (
	"context"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
)

type ProductService struct {
	repository ports.ProductRepository
}

func NewProductService(repository ports.ProductRepository) ports.ProductService {
	return &ProductService{
		repository,
	}
}

// Create implements ports.ProductService.
func (p *ProductService) Create(ctx context.Context, name string) (domain.Product, error) {
	user := ctx.Value("user").(domain.User)

	domainName, err := domain.NewProductName(name)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Create: Error: %w", err)
	}

	product := domain.CreateProduct(domainName, []domain.Category{}, user)

	err = p.repository.Create(ctx, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Create: Error: %w", err)
	}

	created, err := p.Get(ctx, name)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Create: Error: %w", err)
	}

	return created, nil
}

// Get implements ports.ProductService.
func (p *ProductService) Get(ctx context.Context, name string) (domain.Product, error) {
	user := ctx.Value("user").(domain.User)

	domainName, err := domain.NewProductName(name)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Get: Error: %w", err)
	}

	product, err := p.repository.Get(ctx, user, domainName)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Get: Error: %w", err)
	}

	return product, nil
}

// GetAllForUser implements ports.ProductService.
func (p *ProductService) GetAllForUser(ctx context.Context) ([]domain.Product, error) {
	user := ctx.Value("user").(domain.User)

	products, err := p.repository.GetAll(ctx, user)
	if err != nil {
		return []domain.Product{}, fmt.Errorf("services.ProductService.GetAll: Error: %w", err)
	}

	return products, nil
}

// GetById implements ports.ProductService.
func (p *ProductService) GetById(ctx context.Context, id uuid.UUID) (domain.Product, error) {
	user := ctx.Value("user").(domain.User)

  product, err := p.repository.GetById(ctx, user, id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.GetById: Error: %w", err)
	}

	return product, nil
}
