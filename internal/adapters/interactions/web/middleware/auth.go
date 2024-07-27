package middleware

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/auth"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

func AuthMiddleware(userService ports.UserService, loginUrl string, oauthConfig *oauth2.Config) Middleware {
	redirectToLogin := func(writer http.ResponseWriter, request *http.Request) {
		slog.Debug("Redirecting to login")

		if request.Header.Get("HX-Request") != "" {
			writer.Header().Add("HX-Redirect", loginUrl)
			writer.WriteHeader(http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(writer, request, loginUrl, http.StatusTemporaryRedirect)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenCookie, err := request.Cookie(auth.TokenCookieName)
			if err != nil || tokenCookie.Value == "" {
				slog.Debug("Token cookie malformed or missing")
				redirectToLogin(writer, request)
				return
			}

			refreshCookie, err := request.Cookie(auth.RefreshTokenCookieName)
			if err != nil || refreshCookie.Value == "" {
				slog.Debug("Refresh token cookie malformed or missing")
				redirectToLogin(writer, request)
				return
			}

			expiryCookie, err := request.Cookie(auth.TokenExpiryCookieName)
			if err != nil || expiryCookie.Value == "" {
				slog.Debug("Expiry cookie malformed or missing")
				redirectToLogin(writer, request)
				return
			}

			parsedExpiry, err := strconv.Atoi(expiryCookie.Value)
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				io.WriteString(writer, "failed to parse expiry")
				return
			}
			unixTime := time.Unix(int64(parsedExpiry), 0)

			token := &oauth2.Token{
				AccessToken:  tokenCookie.Value,
				RefreshToken: refreshCookie.Value,
				Expiry:       unixTime,
			}

			userinfo, err := auth.GetUserInfo(request.Context(), token, oauthConfig)
			if err != nil {
				slog.Debug("Token not valid for request")
				redirectToLogin(writer, request)
				return
			}

			user, err := userService.Get(request.Context(), userinfo.Email)
			if err != nil {
				slog.Error("failed to get user by email", err)
				writer.WriteHeader(500)
				io.WriteString(writer, "Internal server error")
				return
			}

			if user.Id == uuid.Nil {
				user, err = userService.Create(request.Context(), userinfo.Email)
				if err != nil {
					slog.Error("failed to create user", err)
					writer.WriteHeader(500)
					io.WriteString(writer, "Internal server error")
					return
				}
			}

			slog.Info("User authenticated in successfully")
			request = request.WithContext(context.WithValue(request.Context(), "user", user))

			next.ServeHTTP(writer, request)
		})
	}
}
