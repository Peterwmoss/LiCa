package functions

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/functions/auth"
	"github.com/rs/zerolog/log"
)

func IsHTMXRequest(request *http.Request) bool {
	if request.Header.Get("HX-Request") != "" {
		return true
	}

	return false
}

func GetToken(request *http.Request) (string, error) {
	token, err := request.Cookie("token")
	if err != nil {
		return "", err
	}

	return token.String(), nil
}

func RedirectToLogin(writer http.ResponseWriter, request *http.Request) {
	redirectUrl := auth.BASE_URL + "/login"

	log.Info().Msg("Redirecting to login")

	http.Redirect(writer, request, redirectUrl, http.StatusTemporaryRedirect)
}
