package functions

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
)

type (
	UserHandler struct {
		userService domain.UserService
	}
)

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) Get(writer http.ResponseWriter, request *http.Request) {
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

	Render(writer, "/components/user", user)
}
