package pages

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
)

type Index struct{}

func (h Index) Index(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(domain.User)
	handlers.Templates.Render(writer, "index", user)
}
