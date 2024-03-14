package functions

import (
	"net/http"
	"strconv"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/views"
	"github.com/rs/zerolog/log"
)

type (
	ListHandler struct {
		GetAll http.Handler
		Get    http.Handler
		Create http.Handler
	}

	listHandler struct {
		listService domain.ListService
		templates   *views.Templates
	}
)

func NewListHandler(listService domain.ListService, templates *views.Templates) *ListHandler {
	handler := &listHandler{listService, templates}
	return &ListHandler{
		GetAll: handler.getAll(),
		Get:    handler.get(),
		Create: handler.create(),
	}
}

func (h *listHandler) getAll() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		user := request.Context().Value("user").(domain.User)

		lists, err := h.listService.GetAll(user)
		if err != nil {
			log.Error().Err(err).Msg("failed to get lists for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if IsHTMXRequest(request) {
			h.templates.Render(writer, "lists", lists)
			return
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}

func (h *listHandler) get() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		user := request.Context().Value("user").(domain.User)

		pathId := request.PathValue("id")
		id, err := strconv.Atoi(pathId)
		list, err := h.listService.Get(user, id)
		if err != nil {
			log.Error().Err(err).Msg("failed to get list for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if IsHTMXRequest(request) {
			h.templates.Render(writer, "list", list)
			return
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}

func (h *listHandler) create() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		user := request.Context().Value("user").(*domain.User)

		name := request.FormValue("name")
		lists, err := h.listService.Create(name, *user)
		if err != nil {
			log.Error().Err(err).Msg("failed to create list for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if IsHTMXRequest(request) {
			Render(writer, "/components/lists", lists)
			return
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}
