package pages

import (
	"log/slog"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type List struct {
	ListService ports.ListService
}

func (h List) Lists(writer http.ResponseWriter, request *http.Request) {
	if !handlers.IsHTMXRequest(request) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	lists, err := h.ListService.GetAllForUser(request.Context())
	if err != nil {
		slog.Error("Failed to get lists for user", "error", err)
		handlers.HandleError(writer, err)
		return
	}

	user := request.Context().Value("user").(domain.User)

	data := struct {
		Lists []domain.List
		User  domain.User
	}{
		Lists: lists,
		User:  user,
	}

	slog.Debug("Rendering lists", "data", data)
	err = handlers.Templates.Render(writer, "lists", data)
	if err != nil {
		slog.Error("Failed to render lists", "error", err)
	}
}

func (h List) List(writer http.ResponseWriter, request *http.Request) {
	if !handlers.IsHTMXRequest(request) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	user := request.Context().Value("user").(domain.User)
	listName := request.PathValue("name")

	list, err := h.ListService.Get(request.Context(), listName)
	if err != nil {
		slog.Error("Failed to get list for user", "list name", listName, "error", err)
		handlers.HandleError(writer, err)
		return
	}

	data := struct {
		List domain.List
		User domain.User
	}{
		List: list,
		User: user,
	}

	slog.Debug("Rendering list page", "data", data)
	err = handlers.Templates.Render(writer, "list", data)
	if err != nil {
		slog.Error("Failed to render list page", "error", err)
		handlers.HandleError(writer, err)
	}
}
