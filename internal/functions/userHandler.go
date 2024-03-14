package functions

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/views"
	"github.com/rs/zerolog/log"
)

type (
	UserHandler struct {
		Get http.Handler
	}

	userHandler struct {
		userService domain.UserService
		templates   *views.Templates
	}
)

func NewUserHandler(userService domain.UserService, templates *views.Templates) *UserHandler {
	handler := &userHandler{userService, templates}

	return &UserHandler{
		Get: handler.get(),
	}
}

func (h *userHandler) get() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !IsHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		token, err := GetToken(request)
		if err != nil {
			log.Info().Err(err).Msg("Failed to get token")
			RedirectToLogin(writer, request)
			return
		}

		user, err := h.userService.Get(token)
		if err != nil {
			log.Info().Err(err).Msg("Failed to get user")
			RedirectToLogin(writer, request)
			return
		}

		h.templates.Render(writer, "user", user)
	})
}
