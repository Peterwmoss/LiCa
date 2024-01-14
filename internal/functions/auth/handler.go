package auth

import (
	"context"
	"os"

	"github.com/Peterwmoss/LiCa/internal/auth"
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

const BASE_URL = "/auth"

type (
	Handler interface {
		Mount(*fiber.App)
		Login(*fiber.Ctx) error
		Logout(*fiber.Ctx) error
		Callback(*fiber.Ctx) error
	}

	handler struct {
		authConfig  *oauth2.Config
		userService domain.UserService
		ctx         context.Context
		stateCheck  string
		baseUrl     string
	}
)

func NewHandler(userService domain.UserService, ctx context.Context) Handler {
	stateCheck := os.Getenv("LICA_STATE_CHECK")
	if stateCheck == "" {
		stateCheck = "a8e7hfwnkf3"
	}

	return &handler{
		auth.NewAuthConfig(BASE_URL).Get(),
		userService,
		ctx,
		stateCheck,
		BASE_URL,
	}
}

func (h handler) Mount(app *fiber.App) {
	app.Get(h.baseUrl+"/login", h.Login)
	app.Get(h.baseUrl+"/logout", h.Logout)
	app.Get(h.baseUrl+"/callback", h.Callback)
}

func (h handler) Login(ctx *fiber.Ctx) error {
	redirectUrl := h.authConfig.AuthCodeURL(h.stateCheck)
	return ctx.Status(fiber.StatusSeeOther).Redirect(redirectUrl)
}

func (h handler) Callback(ctx *fiber.Ctx) error {
	state := ctx.Query("state")
	if state != h.stateCheck {
		log.Error().Msg("CSFR detected")
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	code := ctx.Query("code")
	token, err := h.authConfig.Exchange(h.ctx, code)
	if err != nil {
		log.Error().Err(err).Msg("Failed code exchange")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = h.userService.GetOrCreate(token.AccessToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user, see error for more details: ")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: token.AccessToken,
	})
	return ctx.Status(fiber.StatusTemporaryRedirect).Redirect("/")
}

func (h handler) Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: "",
	})
	return ctx.SendStatus(fiber.StatusOK)
}
