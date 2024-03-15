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

func AuthMiddleware(userService domain.UserService, loginUrl string) func(next http.Handler) http.Handler {
	redirectToLogin := func(writer http.ResponseWriter, request *http.Request) {
		log.Debug().Msg("Redirecting to login")
		http.Redirect(writer, request, loginUrl, http.StatusTemporaryRedirect)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenCookie, err := request.Cookie("token")
			if err != nil {
				log.Debug().Msg("Token cookie malformed or missing")
				redirectToLogin(writer, request)
				return
			}

			token := tokenCookie.Value
			if token == "" {
				log.Debug().Msg("Token cookie has no value")
				redirectToLogin(writer, request)
				return
			}

			user, err := userService.GetOrCreate(token)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get user from token")
				redirectToLogin(writer, request)
				return
			}

			log.Info().Msg("User authenticated in successfully")
			request = request.WithContext(context.WithValue(request.Context(), "user", *user))

			next.ServeHTTP(writer, request)
		})
	}
}
