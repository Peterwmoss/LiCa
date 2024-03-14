package functions

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/views"
)

func GetIndex(templates *views.Templates) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		templates.Render(writer, "index", nil)
	})
}
