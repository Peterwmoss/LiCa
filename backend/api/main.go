package main

import (
	"context"

	"github.com/Peterwmoss/LiCa/database"
	"github.com/Peterwmoss/LiCa/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
  ctx := context.Background()

  app := fiber.New()

  db := database.Get()
  defer db.Close()

  database.CreateSchema(ctx)

  err := database.Seed(ctx)
  if err != nil {
    log.Fatal(err)
  }

  app.Get("/items", func (c *fiber.Ctx) error {
    items, err := repository.Items(db, ctx)
    if err != nil {
      log.Warn(err)
    }
    return c.JSON(items)
  })

  app.Get("/categories", func (c *fiber.Ctx) error {
    categories, err := repository.Categories(db, ctx)
    if err != nil {
      log.Warn(err)
    }
    return c.JSON(categories)
  })

  log.Fatal(app.Listen(":3000"))
}
