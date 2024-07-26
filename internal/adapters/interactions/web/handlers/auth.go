package handlers

import (
	"context"
	"crypto/rand"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/auth"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var authStateCheck string

func init() {
	stateCheck, present := os.LookupEnv("LICA_STATE_CHECK")
	if !present {
		bytes := make([]byte, 20)
		rand.Read(bytes)
		stateCheck = fmt.Sprintf("%x", bytes)
		slog.Info("", "State check", stateCheck)
	}
	authStateCheck = stateCheck
}

func AuthLogin(oauth2Config *oauth2.Config) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		redirectUrl := oauth2Config.AuthCodeURL(authStateCheck, oauth2.AccessTypeOffline)
		http.Redirect(writer, request, redirectUrl, http.StatusSeeOther)
	})
}

func AuthCallback(oauth2Config *oauth2.Config, userService ports.UserService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		state := request.URL.Query().Get("state")

		if state != authStateCheck {
			slog.Error("CSFR detected")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		code := request.URL.Query().Get("code")
		token, err := oauth2Config.Exchange(request.Context(), code)
		if err != nil {
			slog.Error("Failed code exchange", "error", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		userinfo, err := auth.GetUserInfo(request.Context(), token, oauth2Config)
		if err != nil {
			slog.Error("Failed to get userinfo", "error", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		email, err := domain.NewEmail(userinfo.Email)
		if err != nil {
			slog.Error("Failed to validate email from userinfo", "error", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		contextWithUser := context.WithValue(request.Context(), "user", domain.NewUser(uuid.Nil, email))
		request = request.WithContext(contextWithUser)

		user, err := userService.Get(request.Context(), string(email))
		if err != nil {
			slog.Error("Failed to get user", "error", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if user.Id == uuid.Nil {
			user, err = userService.Create(request.Context(), string(email))
			if err != nil {
				slog.Error("Failed to create user", "error", err)
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		http.SetCookie(writer, &http.Cookie{
			Name:  auth.TokenCookieName,
			Path:  "/",
			Value: token.AccessToken,
		})

		http.SetCookie(writer, &http.Cookie{
			Name:  auth.TokenExpiryCookieName,
			Path:  "/",
			Value: strconv.Itoa(int(token.Expiry.Unix())),
		})

		http.SetCookie(writer, &http.Cookie{
			Name:  auth.RefreshTokenCookieName,
			Path:  "/",
			Value: token.RefreshToken,
		})

		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
	})
}

func AuthLogout() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.SetCookie(writer, &http.Cookie{
			Name:  auth.TokenCookieName,
			Path:  "/",
      MaxAge: -1,
		})

		http.SetCookie(writer, &http.Cookie{
			Name:  auth.TokenExpiryCookieName,
			Path:  "/",
      MaxAge: -1,
		})

		http.SetCookie(writer, &http.Cookie{
			Name:  auth.RefreshTokenCookieName,
			Path:  "/",
      MaxAge: -1,
		})

		writer.WriteHeader(http.StatusNoContent)
	})
}
