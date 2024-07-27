package actions

import (
	"log/slog"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers/pages"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type List struct {
	ListService ports.ListService
	ListPage pages.List
}

func (l *List) Create(writer http.ResponseWriter, request *http.Request) {
  if !handlers.IsHTMXRequest(request) {
    writer.WriteHeader(http.StatusNotFound)
    return
  }

  name := request.FormValue("name")
  created, err := l.ListService.Create(request.Context(), name)
  if err != nil {
    slog.Error("failed to create list", "error", err)
    handlers.HandleError(writer, err)
    return
  }

  slog.Debug("Created new list for user", "created", created)

  request.SetPathValue("name", name)
  l.ListPage.List(writer, request)
}
