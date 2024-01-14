package list

import (
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/functions"
	"github.com/Peterwmoss/LiCa/internal/middleware"

	"github.com/Peterwmoss/LiCa/internal/templates/partials"
	"github.com/gofiber/fiber/v2"
)

type (
	Handler interface {
		GetAll(*fiber.Ctx) error
		Mount(*fiber.App)
	}

	handler struct {
		listService domain.ListService
		userService domain.UserService
	}
)

func NewHandler(listService domain.ListService, userService domain.UserService) Handler {
	return &handler{listService, userService}
}

func (h handler) Mount(app *fiber.App) {
	authMiddleware := middleware.NewAuth(h.userService)
	app.Use("/lists", authMiddleware.Known)

	app.Get("/lists", h.GetAll)
}

func (h handler) GetAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	lists, err := h.listService.GetAll(*user)
	if err != nil {
		return err
	}

	if functions.IsHTMXRequest(ctx) {
		return functions.ToHandler(partials.Lists(lists))(ctx)
	}

	return ctx.SendStatus(fiber.StatusNotFound)
}
