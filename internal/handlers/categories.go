package handlers

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/rs/zerolog/log"
)

func CategoriesGetAll(categoryService domain.CategoryService) http.Handler {
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

		templates.Render(writer, "categories", categories)
	})
}
