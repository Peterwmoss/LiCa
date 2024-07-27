package components

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type ListItem struct {
	CategoryService ports.CategoryService
}

func (l *ListItem) New(writer http.ResponseWriter, request *http.Request) {
	if !handlers.IsHTMXRequest(request) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	listName := request.URL.Query().Get("list")
	if listName == "" {
		handlers.HandleError(writer, errors.New("no list provided"))
		return
	}

	categories, err := l.CategoryService.GetAll(request.Context())
	if err != nil {
    slog.Error("Failed to get categories", "error", err)
		handlers.HandleError(writer, err)
		return
	}

	data := struct {
		ListName   string
		Categories []domain.Category
	}{
		ListName:   listName,
		Categories: categories,
	}

	handlers.Templates.Render(writer, "list-item-new", data)
}
