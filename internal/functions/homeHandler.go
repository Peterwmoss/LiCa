package functions

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
)

type (
	HomeHandler struct {
		userService domain.UserService
	}
)

func NewHomeHandler(userService domain.UserService) *HomeHandler {
	return &HomeHandler{userService}
}

func (h *HomeHandler) Get(writer http.ResponseWriter, request *http.Request) {
	token, err := GetToken(request)
	if err != nil {
    writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err = h.userService.GetOrCreate(token)
	if err != nil {
		log.Warn().Err(err).Msg("Error while getting user")
    writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	Render(writer, "/pages/home.html", nil)
}
