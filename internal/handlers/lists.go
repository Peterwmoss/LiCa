package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
)

func ListGetAll(listService domain.ListService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		user := request.Context().Value("user").(domain.User)

		lists, err := listService.GetAll(user)
		if err != nil {
			log.Error().Err(err).Msg("failed to get lists for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		templates.Render(writer, "lists", lists)
	})
}

func ListGet(listService domain.ListService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		user := request.Context().Value("user").(domain.User)

		pathId := request.PathValue("id")
		id, err := strconv.Atoi(pathId)
		if err != nil {
			log.Error().Err(err).Msg("invalid id")
			io.WriteString(writer, errors.New("invalid id").Error())
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		list, err := listService.Get(user, id)
		if err != nil {
			log.Error().Err(err).Msg("failed to get list for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		templates.Render(writer, "list", list)
	})
}

func ListCreate(listService domain.ListService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		user := request.Context().Value("user").(domain.User)

		name := request.FormValue("list-name")
		created, err := listService.Create(name, user)

		if errors.Is(err, domain.UniqueViolationError) {
			log.Debug().Msgf("list already exists for user: \"%s\" with name: \"%s\"", user.Email, name)
			writer.WriteHeader(http.StatusConflict)
			io.WriteString(writer, fmt.Sprintf("Liste med navn: \"%s\" eksisterer allerede", name))
			return
		} else if err != nil {
			log.Error().Err(err).Msg("failed to create list for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Debug().Msgf("created list successfully, %d", created.Id)

		templates.Render(writer, "list-item", created)
	})
}

func NewList() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		templates.Render(writer, "new_list", nil)
	})
}
