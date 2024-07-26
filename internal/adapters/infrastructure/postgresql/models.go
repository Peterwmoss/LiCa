package postgresql

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Product struct {
		bun.BaseModel `bun:"table:products,alias:p"`

		Id         uuid.UUID           `bun:",pk"`
		Name       string              `bun:",unique:products_unique,notnull"`
		Categories []ProductCategories `bun:"rel:has-many,join:id=product_id"`
		UserId     uuid.UUID           `bun:",unique:products_unique"`
		User       User                `bun:"rel:has-one,join:user_id=id"`
		IsCustom   bool                `bun:",notnull,default:false"`
	}

	ProductCategories struct {
		bun.BaseModel `bun:"table:product_categories,alias:pc"`

		ProductId  uuid.UUID `bun:",pk"`
		Product    Product   `bun:"rel:has-one,join:product_id=id"`
		CategoryId uuid.UUID `bun:",pk"`
		Category   Category  `bun:"rel:has-one,join:category_id=id"`
		UserId     uuid.UUID `bun:",pk"`
		User       User      `bun:"rel:has-one,join:user_id=id"`
	}

	List struct {
		bun.BaseModel `bun:"table:lists,alias:l"`

		Id        uuid.UUID  `bun:",pk"`
		Name      string     `bun:",notnull,unique:list"`
		UserId    uuid.UUID  `bun:",notnull,unique:list"`
		User      User       `bun:"rel:has-one,join:user_id=id"`
		ListItems []ListItem `bun:"rel:has-many,join:id=list_id"`
	}

	ListItem struct {
		bun.BaseModel `bun:"table:list_items,alias:li"`

		Id         uuid.UUID `bun:",pk"`
		Unit       string    ``
		Amount     float32   `bun:",notnull,default:1.0"`
		ListId     uuid.UUID `bun:",notnull,unique:list_items_unique"`
		List       List      `bun:"rel:has-one,join:list_id=id"`
		ProductId  uuid.UUID `bun:",notnull,unique:list_items_unique"`
		Product    Product   `bun:"rel:has-one,join:product_id=id"`
		CategoryId uuid.UUID `bun:",notnull,unique:list_items_unique"`
		Category   Category  `bun:"rel:has-one,join:category_id=id"`
	}

	User struct {
		bun.BaseModel `bun:"table:users,alias:u"`

		Id    uuid.UUID `bun:",pk"`
		Email string    `bun:",unique,notnull"`
	}

	Category struct {
		bun.BaseModel `bun:"table:categories,alias:c"`

		Id       uuid.UUID           `bun:",pk"`
		Name     string              `bun:",unique:user_category-unique,notnull"`
		UserId   uuid.UUID           `bun:",unique:user_category_unique"`
		User     User                `bun:"rel:has-one,join:user_id=id"`
		Products []ProductCategories `bun:"rel:has-many,join:id=category_id"`
	}
)
