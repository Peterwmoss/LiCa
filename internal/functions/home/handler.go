package home

import (
	"github.com/Peterwmoss/LiCa/internal/functions"
	"github.com/Peterwmoss/LiCa/internal/functions/auth"
	"github.com/Peterwmoss/LiCa/internal/functions/user"
	"github.com/Peterwmoss/LiCa/internal/templates/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type (
	Handler interface {
		Get(*fiber.Ctx) error
		Mount(*fiber.App)
	}

	handler struct {
		userService user.Service
	}
)

func NewHandler(userService user.Service) Handler {
	return &handler{userService}
}

func (h handler) Mount(app *fiber.App) {
	app.Get("/", h.Get)
}

func (h handler) Get(ctx *fiber.Ctx) error {
	token := string(ctx.Request().Header.Cookie("token"))

	user, err := h.userService.Get(token)
	if err != nil {
    log.Warn().Err(err).Msg("Error while getting user, redirecting to login")
		return ctx.Redirect(auth.BASE_URL+"/login", fiber.StatusTemporaryRedirect)
	}

	return functions.ToHandler(pages.HomePage(user))(ctx)
}
