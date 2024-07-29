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
		return domain.Product{}, fmt.Errorf("services.ProductService.Create: failed to create product name: %s\n%w", name, err)
	}

	product := domain.CreateProduct(domainName, []domain.Category{}, user)

	err = p.repository.Create(ctx, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Create: failed to create product: %v\n%w", product, err)
	}

	created, err := p.Get(ctx, name)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Create: failed to get product: %v\n%w", product, err)
	}

	return created, nil
}

// Get implements ports.ProductService.
func (p *ProductService) Get(ctx context.Context, name string) (domain.Product, error) {
	user := ctx.Value("user").(domain.User)

	domainName, err := domain.NewProductName(name)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Get: failed to create product name: %s\n%w", name, err)
	}

	product, err := p.repository.Get(ctx, user, domainName)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.Get: failed to get product with name: %s\n%w", domainName, err)
	}

	return product, nil
}

// GetAllForUser implements ports.ProductService.
func (p *ProductService) GetAllForUser(ctx context.Context) ([]domain.Product, error) {
	user := ctx.Value("user").(domain.User)

	products, err := p.repository.GetAll(ctx, user)
	if err != nil {
		return []domain.Product{}, fmt.Errorf("services.ProductService.GetAll: failed to get products for user: %v\n%w", user, err)
	}

	return products, nil
}

// GetById implements ports.ProductService.
func (p *ProductService) GetById(ctx context.Context, id uuid.UUID) (domain.Product, error) {
	user := ctx.Value("user").(domain.User)

  product, err := p.repository.GetById(ctx, user, id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("services.ProductService.GetById: failed to get product with id: %s\n%w", id, err)
	}

	return product, nil
}
