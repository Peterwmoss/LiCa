package handlers

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/views"
)

var templates *views.Templates

func init() {
	templates = views.NewTemplates()
}

func isHTMXRequest(request *http.Request) bool {
	if request.Header.Get("HX-Request") != "" {
		return true
	}

	return false
}
