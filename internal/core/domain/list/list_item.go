package list

import (
	"github.com/Peterwmoss/LiCa/internal/core/domain/category"
	"github.com/Peterwmoss/LiCa/internal/core/domain/product"
	"github.com/google/uuid"
)

type ListItem struct {
	Id             uuid.UUID
	Product        *product.Product
	Amount         Amount
	Unit           Unit
	CustomCategory *category.Category
}
