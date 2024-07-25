package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type ListHandler struct{
  listService ports.ListService
}

func (h ListHandler) GetAll() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		lists, err := h.listService.GetAllForUser(request.Context())
		if err != nil {
			slog.Error("failed to get lists for user", "error", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		templates.Render(writer, "lists", lists)
	})
}

func (h ListHandler) Get() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		listName := request.PathValue("name")

		list, err := h.listService.Get(request.Context(), listName)
		if err != nil {
			slog.Error("failed to get list for user", "error", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		templates.Render(writer, "list", list)
	})
}

func (h ListHandler) Create() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		name := request.FormValue("list-name")
		created, err := h.listService.Create(request.Context(), name)
    if err != nil {
			slog.Error("failed to create list", "error", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
    }

		slog.Debug("created list successfully", "id", created.Id)

		templates.Render(writer, "list-item", created)
	})
}
