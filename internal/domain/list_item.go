package domain

import "github.com/Peterwmoss/LiCa/internal/database"

type (
	ListItem struct {
		Id      int
		Product Product
		Amount  float32
		Unit    Unit
	}

	ListItemService interface {
		ToDomain(database.ListItem) ListItem
		ToDatabase(ListItem, List) database.ListItem
	}
	listItemService struct {
		productService ProductService
	}
)

func NewListItemService(productService ProductService) ListItemService {
	return &listItemService{productService}
}

func (svc listItemService) ToDomain(listItem database.ListItem) ListItem {
	return ListItem{
		Id:      listItem.Id,
		Product: svc.productService.ToDomain(*listItem.Product),
		Amount:  float32(listItem.Amount),
		Unit:    Unit{*listItem.Unit},
	}
}

func (svc listItemService) ToDatabase(listItem ListItem, list List) database.ListItem {
	return database.ListItem{
		Id:        listItem.Id,
		Unit:      &listItem.Unit.Unit,
		Amount:    listItem.Amount,
		ListId:    list.Id,
		ProductId: listItem.Product.id,
	}
}
