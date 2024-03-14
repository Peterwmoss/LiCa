package functions

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
)

type (
	ListHandler struct {
		listService domain.ListService
		userService domain.UserService
	}
)

func NewListHandler(listService domain.ListService, userService domain.UserService) *ListHandler {
	return &ListHandler{listService, userService}
}

func (h *ListHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(*domain.User)

	lists, err := h.listService.GetAll(*user)
	if err != nil {
		log.Error().Err(err).Msg("failed to get lists for user")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if IsHTMXRequest(request) {
		Render(writer, "/components/lists", lists)
		return
	}

	writer.WriteHeader(http.StatusNotFound)
}

func (h *ListHandler) Create(writer http.ResponseWriter, request *http.Request) {
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
}
