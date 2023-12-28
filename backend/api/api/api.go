package api

import (
	"context"
	"fmt"

	"github.com/Peterwmoss/LiCa/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start(port int16, ctx context.Context) error {
	app := fiber.New()

	db := database.Get()
	defer db.Close()

  err := database.CreateSchema(ctx)
  if err != nil {
    return err
  }

  err = database.Seed(db, ctx)
  if err != nil {
    return err
  }

  authCfg := getAuthConfig()

	app.Use(cors.New())

  app.Get("/items", items(db, ctx))
  app.Get("/categories", categories(db, ctx))

	app.Get("/auth/login", loginHandler(authCfg))
	app.Get("/auth/logout", logoutHandler())
	app.Get("/auth/callback", callbackHandler(authCfg, db, ctx))
  app.Get("/auth/me", meHandler(db, ctx))

	return app.Listen(":" + fmt.Sprintf("%d", port))
}
