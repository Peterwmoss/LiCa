package pages

import (
	"log/slog"
	"net/http"
	"slices"

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

	slog.Debug("Rendering lists", "data", lists)
	err = handlers.Templates.Render(writer, "lists", lists)
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

	type DataCategories struct {
		Items    []domain.ListItem
		Category domain.Category
	}

	dataCategories := []DataCategories{}

	for _, item := range list.Items {
		containedIdx := slices.IndexFunc(dataCategories, func(dataCat DataCategories) bool {
			return item.Category.Id == dataCat.Category.Id
		})
    
		if containedIdx != -1 {
			dataCat := &dataCategories[containedIdx]
			dataCat.Items = append(dataCat.Items, item)
			continue
		}

		dataCategories = append(dataCategories, DataCategories{
			Items:    []domain.ListItem{item},
			Category: item.Category,
		})
	}

	data := struct {
		Name       domain.ListName
		User       domain.User
		Categories []DataCategories
	}{
		Name:       list.Name,
		User:       user,
		Categories: dataCategories,
	}

	slog.Debug("Rendering list page", "data", data)
	err = handlers.Templates.Render(writer, "list", data)
	if err != nil {
		slog.Error("Failed to render list page", "error", err)
		handlers.HandleError(writer, err)
	}
}
