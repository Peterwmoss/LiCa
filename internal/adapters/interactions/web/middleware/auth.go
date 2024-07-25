package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/auth"
	web "github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/errors"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
)

func AuthMiddleware(userService ports.UserService, loginUrl string) Middleware {
	redirectToLogin := func(writer http.ResponseWriter, request *http.Request) {
		slog.Debug("Redirecting to login")
		http.Redirect(writer, request, loginUrl, http.StatusTemporaryRedirect)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenCookie, err := request.Cookie("token")
			if err != nil {
				slog.Debug("Token cookie malformed or missing")
				redirectToLogin(writer, request)
				return
			}

			token := tokenCookie.Value
			if token == "" {
				slog.Debug("Token cookie has no value")
				redirectToLogin(writer, request)
				return
			}

			userinfo, err := auth.GetUserInfo(token)
			if err != nil {
				slog.Debug("Token not valid for request")
				redirectToLogin(writer, request)
				return
			}

			user, err := userService.Get(request.Context(), userinfo.Email)
			if err != nil {
				slog.Error("failed to get user by email", err)
				writer.WriteHeader(500)
				return
			}

			if user.Id == uuid.Nil {
				user, err = userService.Create(request.Context(), userinfo.Email)
				if err != nil {
					slog.Error("failed to create user", err)
          web.HandleError(err, writer)
					return
				}
			}

			slog.Info("User authenticated in successfully")
			request = request.WithContext(context.WithValue(request.Context(), "user", user))

			next.ServeHTTP(writer, request)
		})
	}
}
