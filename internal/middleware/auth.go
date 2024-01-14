package middleware

import (
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type (
	Auth interface {
		Known(*fiber.Ctx) error
	}

	auth struct {
		userService domain.UserService
	}
)

func NewAuth(userService domain.UserService) Auth {
	return &auth{userService}
}

func (a auth) Known(ctx *fiber.Ctx) error {
	token := string(ctx.Request().Header.Cookie("token"))
	if token == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	user, err := a.userService.Get(token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user from token")
		ctx.SendStatus(fiber.StatusUnauthorized)
	}

	ctx.Locals("user", user)

	return ctx.Next()
}
