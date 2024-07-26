package components

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers"
)

type List struct{}

func (l *List) New(writer http.ResponseWriter, request *http.Request) {
	if !handlers.IsHTMXRequest(request) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	handlers.Templates.Render(writer, "list-new", nil)
}
