package home

import (
	"context"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/functions"
	"github.com/Peterwmoss/LiCa/internal/templates/pages"
	"github.com/rs/zerolog/log"
)

type (
	handler struct {
		userService domain.UserService
	}
)

func NewHandler(userService domain.UserService) http.Handler {
	server := http.NewServeMux()

	homeHandler := &handler{userService}

	server.HandleFunc("GET /", homeHandler.get)

	return server
}

func (h *handler) get(writer http.ResponseWriter, request *http.Request) {
	token, err := functions.GetToken(request)
	if err != nil {
		functions.RedirectToLogin(writer, request)
		return
	}

	user, err := h.userService.Get(token)
	if err != nil {
		log.Warn().Err(err).Msg("Error while getting user, redirecting to login")
		functions.RedirectToLogin(writer, request)
		return
	}

	pages.HomePage(*user).Render(context.Background(), writer)
}
