package user

import (
	"context"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/functions"
	"github.com/Peterwmoss/LiCa/internal/templates/partials"
	"github.com/rs/zerolog/log"
)

func NewHandler(userService domain.UserService) http.Handler {
	server := http.NewServeMux()

	server.HandleFunc("GET /", get(userService))

	return server
}

func get(userService domain.UserService) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !functions.IsHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		token, err := functions.GetToken(request)
		if err != nil {
			log.Info().Err(err).Msg("Failed to get token")
			functions.RedirectToLogin(writer, request)
			return
		}

		user, err := userService.Get(token)
		if err != nil {
			log.Info().Err(err).Msg("Failed to get user")
			functions.RedirectToLogin(writer, request)
			return
		}

		partials.User(*user).Render(context.Background(), writer)
	}
}
