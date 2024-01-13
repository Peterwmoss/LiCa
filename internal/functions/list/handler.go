package list

import (
	"context"

	"github.com/Peterwmoss/LiCa/internal/functions"
	"github.com/Peterwmoss/LiCa/internal/repository"
	"github.com/Peterwmoss/LiCa/internal/templates/partials"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type (
  Handler interface {
    GetAll(*fiber.Ctx) error
    Mount(*fiber.App)
  }

  handler struct {
    db *bun.DB
    ctx context.Context
  }
)

func NewHandler(db *bun.DB, ctx context.Context) Handler {
  return &handler{ db, ctx }
}

func (h handler) Mount(app *fiber.App) {
  app.Get("/list", h.GetAll)
}

func (h handler) GetAll(ctx *fiber.Ctx) error {
  items, err := repository.GetItems(h.db, h.ctx)
  if err != nil {
    return err
  }

  if functions.IsHTMXRequest(ctx) {
    return functions.ToHandler(partials.Items(items))(ctx)
  }

  return ctx.SendStatus(fiber.StatusNotFound)
}
