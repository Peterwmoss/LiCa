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

func UseAuth(userService domain.UserService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
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

		user, err := userService.GetOrCreate(token)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get user from token")
			redirectToLogin(writer, request)
			return
		}

    log.Info().Msgf("Email: %s", user.Email)
		request = request.WithContext(context.WithValue(request.Context(), "user", *user))

		next.ServeHTTP(writer, request)
	})
}

func redirectToLogin(writer http.ResponseWriter, request *http.Request) {
	redirectUrl := "/auth/login"

	log.Info().Msg("Redirecting to login")

	http.Redirect(writer, request, redirectUrl, http.StatusTemporaryRedirect)
}
