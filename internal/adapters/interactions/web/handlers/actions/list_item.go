package actions

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers/pages"
	"github.com/Peterwmoss/LiCa/internal/core"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type ListItem struct {
	ListItemService ports.ListItemService
	ListPage        pages.List
}

func (l *ListItem) Add(writer http.ResponseWriter, request *http.Request) {
	if !handlers.IsHTMXRequest(request) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	var validationErrors error
	product := request.FormValue("product")
	if product == "" {
		validationErrors = errors.Join(validationErrors, errors.New("product must not be empty"))
	}

	category := request.FormValue("category")
	if category == "" {
		validationErrors = errors.Join(validationErrors, errors.New("category must not be empty"))
	}

	amount, err := strconv.ParseFloat(request.FormValue("amount"), 32)
	if err != nil {
		validationErrors = errors.Join(validationErrors, err)
	}

	if validationErrors != nil {
		slog.Error("Failed to validate input", "error", validationErrors)
		handlers.HandleError(writer, errors.Join(core.ErrValidation, validationErrors))
		return
	}

	listName := request.FormValue("listName")

	item := ports.ListItemCreate{
		ListName:     listName,
		ProductName:  product,
		CategoryName: category,
		Amount:       float32(amount),
	}
	created, err := l.ListItemService.Add(request.Context(), item)
	if err != nil {
		slog.Error("failed to add item to list", "error", err)
		handlers.HandleError(writer, err)
		return
	}

	slog.Debug("Added item to list", "created", created)

	request.SetPathValue("name", listName)
	l.ListPage.List(writer, request)
}
