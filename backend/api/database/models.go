package database

import (
	"fmt"

	"github.com/uptrace/bun"
)

type Item struct {
	bun.BaseModel `bun:"table:items,alias:i"`

	Id         int64     `bun:",pk,autoincrement" json:"id"`
	Name       string    `bun:",unique,notnull" json:"name,omitempty"`
	CategoryId int64     `bun:",notnull" json:"-"`
	Category   *Category `bun:"rel:has-one,join:category_id=id" json:"category,omitempty"`
  UserId     *int64     `json:"-"`
	User       *User     `bun:"rel:has-one,join:user_id=id" json:"user,omitempty"`
}

func (item *Item) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Category: %s, User: %s", item.Id, item.Name, item.Category.String(), item.User.String())
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

  Id   int64  `bun:",pk,autoincrement" json:"id"`
	Name string `bun:",unique,notnull" json:"name"`
}

func (user *User) String() string {
	return fmt.Sprintf("Id: %d, Name: %s", user.Id, user.Name)
}

type Category struct {
	bun.BaseModel `bun:"table:categories,alias:c"`

	Id   int64  `bun:",pk,autoincrement" json:"id"`
  Name string `bun:",unique,notnull" json:"name"`
}

func (category *Category) String() string {
	return fmt.Sprintf("Id: %d, Name: %s", category.Id, category.Name)
}
