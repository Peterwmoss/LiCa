package domain

import "github.com/Peterwmoss/LiCa/internal/database"

type (
	ListItem struct {
		id      int
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
		id:      listItem.Id,
		Product: svc.productService.ToDomain(*listItem.Product),
		Amount:  float32(listItem.Amount),
		Unit:    Unit{*listItem.Unit},
	}
}

func (svc listItemService) ToDatabase(listItem ListItem, list List) database.ListItem {
	return database.ListItem{
		Id:        listItem.id,
		Unit:      &listItem.Unit.Unit,
		Amount:    listItem.Amount,
		ListId:    list.id,
		ProductId: listItem.Product.id,
	}
}
