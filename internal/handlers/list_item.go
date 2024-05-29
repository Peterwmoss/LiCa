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

func ListItemCreate(productService domain.ProductService, categoryService domain.CategoryService, listService domain.ListService, listItemService domain.ListItemService) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !isHTMXRequest(request) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		user := request.Context().Value("user").(domain.User)

		categoryId, err := strconv.Atoi(request.FormValue("category-id"))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			io.WriteString(writer, "Kategori-ID skal være et heltal")
			return
		}

		category, err := categoryService.GetById(categoryId)
		if err != nil {
			// TODO: handle
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		productName := request.FormValue("item-name")
		product, err := productService.CreateIfNotExists(productName, category, user)
		if err != nil {
			// TODO: handle
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

    listId, err := strconv.Atoi(request.PathValue("id"))
    if err != nil {
			// TODO: handle
			writer.WriteHeader(http.StatusInternalServerError)
			return
    }

    list, err := listService.Get(user, listId)
    if err != nil {
			// TODO: handle
			writer.WriteHeader(http.StatusInternalServerError)
			return
    }

		created, err := listItemService.Create(product, list)

		if errors.Is(err, domain.UniqueViolationError) {
			log.Debug().Msgf("list already exists for user: \"%s\" with name: \"%s\"", user.Email, productName)
			writer.WriteHeader(http.StatusConflict)
			io.WriteString(writer, fmt.Sprintf("Liste med navn: \"%s\" eksisterer allerede", productName))
			return
		}

		if errors.Is(err, domain.EmptyNameError) {
			log.Debug().Msgf("user: \"%s\" tried to create list with empty name", user.Email)
			writer.WriteHeader(http.StatusBadRequest)
			io.WriteString(writer, "En liste skal have et navn")
			return
		}

		if err != nil {
			log.Error().Err(err).Msg("failed to create list for user")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Debug().Msgf("created list successfully, %d", created.Id)

		templates.Render(writer, "list-item", created)
	})
}
