package handlers

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/functions"
	"github.com/rs/zerolog/log"
)

func ListGetAll(listService domain.ListService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		user := request.Context().Value("user").(domain.User)

		lists, err := listService.GetAll(user)
		if err != nil {
			log.Error().Err(err).Msg("failed to get lists for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if functions.IsHTMXRequest(request) {
			templates.Render(writer, "lists", lists)
			return
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}

func ListGet(listService domain.ListService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
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

		if functions.IsHTMXRequest(request) {
			templates.Render(writer, "list", list)
			return
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}

func ListCreate(listService domain.ListService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		user := request.Context().Value("user").(*domain.User)

		name := request.FormValue("name")
		lists, err := listService.Create(name, *user)
		if err != nil {
			log.Error().Err(err).Msg("failed to create list for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if functions.IsHTMXRequest(request) {
			templates.Render(writer, "/components/lists", lists)
			return
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}
