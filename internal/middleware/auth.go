package middleware

import (
	"context"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
)

type (
	auth struct {
		userService domain.UserService
	}
)

func NewAuth(userService domain.UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tokenCookie, err := request.Cookie("token")
		if err != nil {
      redirectToLogin(writer, request)
			return
		}

		token := tokenCookie.Value
		if token == "" {
      redirectToLogin(writer, request)
			return
		}

		user, err := userService.Get(token)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get user from token")
      redirectToLogin(writer, request)
			return
		}

		request.WithContext(context.WithValue(request.Context(), "user", user))
	}
}

func redirectToLogin(writer http.ResponseWriter, request *http.Request) {
	redirectUrl := "/auth/login"

	log.Info().Msg("Redirecting to login")

	http.Redirect(writer, request, redirectUrl, http.StatusTemporaryRedirect)
}
