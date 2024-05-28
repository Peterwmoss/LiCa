package middleware

import (
	"context"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/auth"
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
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

			userinfo, err := auth.GetUserInfo(token)
			if err != nil {
				log.Debug().Msg("Token not valid for request")
				redirectToLogin(writer, request)
				return
			}

			user, err := userService.GetOrCreate(userinfo.Email)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get user from email")
				redirectToLogin(writer, request)
				return
			}

			log.Info().Msg("User authenticated in successfully")
			request = request.WithContext(context.WithValue(request.Context(), "user", *user))

			next.ServeHTTP(writer, request)
		})
	}
}
