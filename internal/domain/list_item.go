package domain

import (
	"context"
	"errors"

	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/jackc/pgerrcode"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

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
		Create(*Product, *List) (*ListItem, error)
	}
	listItemService struct {
		productService  ProductService
		categoryService CategoryService
		db              *bun.DB
		ctx             context.Context
	}
)

func NewListItemService(db *bun.DB, ctx context.Context, productService ProductService, categoryService CategoryService) ListItemService {
	return &listItemService{productService, categoryService, db, ctx}
}

func (svc listItemService) Create(product *Product, list *List) (*ListItem, error) {
	listItem := ListItem{
		Product: *product,
		Amount:  1,
		Unit:    Unit{Unit: "stk."},
	}

	databaseListItem := svc.ToDatabase(listItem, *list)

	_, err := svc.db.NewInsert().
		Model(&databaseListItem).
		Exec(svc.ctx)
	if err != nil {
		errStatusCode := err.(pgdriver.Error).Field('C')
		if errStatusCode == pgerrcode.UniqueViolation {
			return nil, errors.Join(err, UniqueViolationError)
		}

		return nil, err
	}

	return &listItem, nil
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
		ProductId: listItem.Product.Id,
	}
}
