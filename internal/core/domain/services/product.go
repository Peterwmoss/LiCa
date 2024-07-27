package services

import (
	"context"
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
		return domain.Product{}, err
	}

	product := domain.CreateProduct(domainName, []domain.Category{}, user)

	err = p.repository.Create(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}

	return p.Get(ctx, name)
}

// Get implements ports.ProductService.
func (p *ProductService) Get(ctx context.Context, name string) (domain.Product, error) {
	user := ctx.Value("user").(domain.User)

	domainName, err := domain.NewProductName(name)
	if err != nil {
		return domain.Product{}, err
	}

	return p.repository.Get(ctx, user, domainName)
}

// GetAllForUser implements ports.ProductService.
func (p *ProductService) GetAllForUser(ctx context.Context) ([]domain.Product, error) {
	user := ctx.Value("user").(domain.User)

	return p.repository.GetAll(ctx, user)
}

// GetById implements ports.ProductService.
func (p *ProductService) GetById(ctx context.Context, id uuid.UUID) (domain.Product, error) {
	user := ctx.Value("user").(domain.User)

	return p.repository.GetById(ctx, user, id)
}
