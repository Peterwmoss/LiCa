package handlers

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

func AuthLogin(oauth2Config *oauth2.Config, stateCheck string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		redirectUrl := oauth2Config.AuthCodeURL(stateCheck)
		http.Redirect(writer, request, redirectUrl, http.StatusSeeOther)
	})
}

func AuthCallback(oauth2Config *oauth2.Config, stateCheck string, userService domain.UserService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		state := request.URL.Query().Get("state")

		if state != stateCheck {
			log.Error().Msg("CSFR detected")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		code := request.URL.Query().Get("code")
		token, err := oauth2Config.Exchange(request.Context(), code)
		if err != nil {
			log.Error().Err(err).Msg("Failed code exchange")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = userService.GetOrCreate(token.AccessToken)
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

func AuthLogout() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.SetCookie(writer, &http.Cookie{
			Name:  "token",
			Value: "",
		})

		writer.WriteHeader(http.StatusNoContent)
	})
}
