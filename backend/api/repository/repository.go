package repository

import (
	"context"

	"github.com/Peterwmoss/LiCa/database"
	"github.com/uptrace/bun"
)

func Items(db *bun.DB, ctx context.Context) ([]*database.Item, error) {
  items := []*database.Item{}

  err := db.NewSelect().Model(&items).Relation("Category").Scan(ctx)
  if err != nil {
    return nil, err
  }

  return items, nil
}

func Categories(db *bun.DB, ctx context.Context) ([]*database.Category, error) {
  items := []*database.Category{}

  err := db.NewSelect().Model(&items).Scan(ctx)
  if err != nil {
    return nil, err
  }

  return items, nil
}
