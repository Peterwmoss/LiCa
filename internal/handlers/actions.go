package handlers

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
)

func NewList() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		templates.Render(writer, "new_list", nil)
	})
}

func NewItem() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

    listId := request.PathValue("id")

		templates.Render(writer, "new_item", listId)
	})
}
func SelectCategory(categoryService domain.CategoryService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		user := request.Context().Value("user").(domain.User)

		categories, err := categoryService.GetAll(user)
		if err != nil {
			log.Error().Err(err).Msg("failed to get categories for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		templates.Render(writer, "category-options", categories)
	})
}
