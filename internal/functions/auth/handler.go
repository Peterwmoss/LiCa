package auth

import (
	"context"
	"os"

	"github.com/Peterwmoss/LiCa/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"golang.org/x/oauth2"
)

const BASE_URL = "/auth"

type (
	Handler interface {
		Mount(*fiber.App)
		Login(*fiber.Ctx) error
		Logout(*fiber.Ctx) error
		Callback(*fiber.Ctx) error
		Info(*fiber.Ctx) error
	}

	handler struct {
		authConfig *oauth2.Config
		db         *bun.DB
		ctx        context.Context
		stateCheck string
		baseUrl    string
	}
)

func NewHandler(db *bun.DB, ctx context.Context) Handler {
	stateCheck := os.Getenv("LICA_STATE_CHECK")
	if stateCheck == "" {
		stateCheck = "a8e7hfwnkf3"
	}

	return &handler{
		authConfig{}.Get(),
		db,
		ctx,
		stateCheck,
		BASE_URL,
	}
}

func (h handler) Mount(app *fiber.App) {
	app.Get(h.baseUrl+"/login", h.Login)
	app.Get(h.baseUrl+"/logout", h.Logout)
	app.Get(h.baseUrl+"/callback", h.Callback)
	app.Get(h.baseUrl+"/info", h.Info)
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

	info, err := GetUserInfo(token.AccessToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user info from Google")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = repository.CreateUser(info.Email, h.db, h.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user")
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

func (h handler) Info(ctx *fiber.Ctx) error {
	token := ctx.Cookies("token")
	if token == "" {
		return ctx.Status(fiber.StatusTemporaryRedirect).Redirect("/auth/login")
	}

	info, err := GetUserInfo(token)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	user, err := repository.GetUser(info.Email, h.db, h.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(user)
}
