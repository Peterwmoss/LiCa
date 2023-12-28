package api

import (
  "context"

	"github.com/Peterwmoss/LiCa/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/uptrace/bun"
)

func items(db *bun.DB, ctx context.Context) (func (c *fiber.Ctx) error) {
	return func(c *fiber.Ctx) error {
		items, err := repository.GetItems(db, ctx)
		if err != nil {
			log.Warn(err)
		}
		return c.JSON(items)
	}
}

func categories(db *bun.DB, ctx context.Context) (func (c *fiber.Ctx) error) {
	return func(c *fiber.Ctx) error {
		categories, err := repository.GetCategories(db, ctx)
		if err != nil {
			log.Warn(err)
		}
		return c.JSON(categories)
	}
}
