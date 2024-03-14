package functions

import (
	"net/http"
	"os"

	"github.com/Peterwmoss/LiCa/internal/auth"
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

var AuthBaseUrl string

type (
	AuthHandler struct {
    Login http.Handler
    Logout http.Handler
    Callback http.Handler
	}

  authHandler struct {
		authConfig  *oauth2.Config
		userService domain.UserService
		stateCheck  string
  }
)

func NewAuthHandler(userService domain.UserService, baseUrl string) *AuthHandler {
	AuthBaseUrl = baseUrl

	stateCheck := os.Getenv("LICA_STATE_CHECK")
	if stateCheck == "" {
		stateCheck = "a8e7hfwnkf3"
	}

  handler := &authHandler{
		authConfig:  auth.NewAuthConfig(baseUrl).Get(),
		userService: userService,
		stateCheck:  stateCheck,
  }

	return &AuthHandler{
    Login: handler.login(),
    Logout: handler.logout(),
    Callback: handler.callback(),
	}
}

func (h *authHandler) login() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		redirectUrl := h.authConfig.AuthCodeURL(h.stateCheck)
		http.Redirect(writer, request, redirectUrl, http.StatusSeeOther)
	})
}

func (h *authHandler) callback() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		state := request.URL.Query().Get("state")
		if state != h.stateCheck {
			log.Error().Msg("CSFR detected")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		code := request.URL.Query().Get("code")
		token, err := h.authConfig.Exchange(request.Context(), code)
		if err != nil {
			log.Error().Err(err).Msg("Failed code exchange")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = h.userService.GetOrCreate(token.AccessToken)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get user, see error for more details: ")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(writer, &http.Cookie{
			Name:  "token",
			Path:  "/",
			Value: token.AccessToken,
		})

		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
	})
}

func (h *authHandler) logout() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.SetCookie(writer, &http.Cookie{
			Name:  "token",
			Value: "",
		})

		writer.WriteHeader(http.StatusOK)

	})
}
