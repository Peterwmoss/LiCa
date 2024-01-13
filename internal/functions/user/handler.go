package user

import (
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/functions"
	"github.com/Peterwmoss/LiCa/internal/templates/partials"
	"github.com/gofiber/fiber/v2"
)

type (
  Handler interface {
    Mount(*fiber.App)
    Get(*fiber.Ctx) error
  }

  handler struct { 
    userService domain.UserService
  }
)

func NewHandler() Handler {
  return &handler{}
}

func (h handler) Mount(app *fiber.App) {
  app.Get("/user", h.Get)
}

func (h handler) Get(ctx *fiber.Ctx) error {
  if !functions.IsHTMXRequest(ctx) {
    return ctx.SendStatus(fiber.StatusNotFound)
  }

  token := string(ctx.Request().Header.Cookie("token"))
  user, err := h.userService.Get(token)
  if err != nil {
    return err
  }

  return functions.ToHandler(partials.User(user))(ctx)
}
