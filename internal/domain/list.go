package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/jackc/pgerrcode"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type (
	List struct {
		Id    int
		Name  string
		Items []ListItem
	}

	ListService interface {
		GetAll(User) ([]List, error)
		Get(User, int) (*List, error)
		Create(name string, user User) (*List, error)
		ToDomain(database.List) List
	}

	listService struct {
		db              *bun.DB
		ctx             context.Context
		listItemService ListItemService
	}
)

func (l List) String() string {
	return fmt.Sprintf("Name: %s", l.Name)
}

func NewListService(db *bun.DB, ctx context.Context, listItemService ListItemService) ListService {
	return &listService{db, ctx, listItemService}
}

func (svc listService) GetAll(user User) ([]List, error) {
	dbLists := []database.List{}

	err := svc.db.NewSelect().
		Model(&dbLists).
		Where("user_id = ?", user.id).
		Scan(svc.ctx)
	if err != nil {
		return nil, err
	}

	lists := make([]List, len(dbLists))
	for i := range dbLists {
		lists[i] = svc.ToDomain(dbLists[i])
	}

	log.Info().Msgf("Found lists: %v", lists)

	return lists, nil
}

func (svc listService) Get(user User, id int) (*List, error) {
	dbList := database.List{}

	err := svc.db.NewSelect().
		Model(&dbList).
		Where("id = ?", id).
		Where("user_id = ?", user.id).
		Relation("ListItems").
		Relation("ListItems.Product").
		Relation("ListItems.Product.Category").
		Relation("ListItems.Product.Category.Orders").
		Limit(1).
		Scan(svc.ctx)
	if err != nil {
		return nil, err
	}

	list := svc.ToDomain(dbList)

	log.Info().Msgf("Found list: %v", list)

	return &list, nil
}

func (svc listService) Create(name string, user User) (*List, error) {
  if name == "" {
    return nil, EmptyNameError
  }

	list := database.List{
		Name:   name,
		UserId: user.id,
	}

	_, err := svc.db.NewInsert().
		Model(&list).
		Exec(svc.ctx)
	if err != nil {
    errStatusCode := err.(pgdriver.Error).Field('C')
		if errStatusCode == pgerrcode.UniqueViolation {
      return nil, errors.Join(err, UniqueViolationError)
		}

		return nil, err
	}

	return svc.Get(user, list.Id)
}

func (svc listService) ToDomain(list database.List) List {
	listItems := make([]ListItem, len(list.ListItems))
	for i := range list.ListItems {
		listItems[i] = svc.listItemService.ToDomain(*list.ListItems[i])
	}

	return List{
		Id:    list.Id,
		Name:  list.Name,
		Items: listItems,
	}
}
