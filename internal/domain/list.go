package domain

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/uptrace/bun"
)

type (
	List struct {
		id    int64
		Items []ListItem
	}

	ListService interface {
		GetAll(User) ([]List, error)
		ToDomain(database.List) List
	}

	listService struct {
		db              *bun.DB
		ctx             context.Context
		listItemService ListItemService
	}
)

func NewListService(db *bun.DB, ctx context.Context, listItemService ListItemService) ListService {
	return &listService{db, ctx, listItemService}
}

func (svc listService) GetAll(user User) ([]List, error) {
	dbLists := []database.List{}

	err := svc.db.NewSelect().
		Model(&dbLists).
		Where("user_id = ?", user.id).
		Relation("ListItems").
		Relation("ListItems.Product").
		Relation("ListItems.Product.Category").
		Relation("ListItems.Product.Category.Orders").
		Scan(svc.ctx)
	if err != nil {
		return nil, err
	}

	lists := make([]List, len(dbLists))
	for i := range dbLists {
		lists[i] = svc.ToDomain(dbLists[i])
	}

	return lists, nil
}

func (svc listService) ToDomain(list database.List) List {
	listItems := make([]ListItem, len(list.ListItems))
	for i := range list.ListItems {
		listItems[i] = svc.listItemService.ToDomain(*list.ListItems[i])
	}

	return List{
		id:    list.Id,
		Items: listItems,
	}
}
