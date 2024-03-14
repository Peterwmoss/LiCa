package functions

import (
	"fmt"
	"html/template"
	"net/http"

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
	redirectUrl := AuthBaseUrl + "/login"

	log.Info().Msg("Redirecting to login")

	http.Redirect(writer, request, redirectUrl, http.StatusTemporaryRedirect)
}

func Render(writer http.ResponseWriter, filePath string, data any) {
  tmpl, err := template.ParseFiles("./internal/templates" + filePath)
  if err != nil {
    log.Error().Err(err).Msg(fmt.Sprintf("Failed to render %s", filePath))
    writer.WriteHeader(http.StatusInternalServerError)
  }

  tmpl.Execute(writer, data)
}
