package repository

import (
	"context"

	"github.com/Peterwmoss/LiCa/database"
	"github.com/uptrace/bun"
)

func GetItems(db *bun.DB, ctx context.Context) ([]*database.Item, error) {
  items := []*database.Item{}

  err := db.NewSelect().Model(&items).Relation("Category").Scan(ctx)
  if err != nil {
    return nil, err
  }

  return items, nil
}

func GetCategories(db *bun.DB, ctx context.Context) ([]*database.Category, error) {
  items := []*database.Category{}

  err := db.NewSelect().Model(&items).Scan(ctx)
  if err != nil {
    return nil, err
  }

  return items, nil
}

func GetUser(email string, db *bun.DB, ctx context.Context) (*database.User, error) {
  user := database.User{}

  err := db.NewSelect().Model(&user).Where("email = ?", email).Limit(1).Scan(ctx)
  if err != nil {
    return nil, err
  }

  return &user, nil
}

func CreateUser(email string, db *bun.DB, ctx context.Context) (*database.User, error) {
  user := database.User{
    Email: email,
  }

  _, err := db.NewInsert().Model(&user).On("CONFLICT DO NOTHING").Exec(ctx)
  if err != nil {
    return nil, err
  }

  return GetUser(email, db, ctx)
}
