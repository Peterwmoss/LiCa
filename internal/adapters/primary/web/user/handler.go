package user

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	web "github.com/Peterwmoss/LiCa/internal/adapters/primary/web/errors"
	"github.com/Peterwmoss/LiCa/internal/core/ports/user"
)

type Handler struct {
	port ports.UserPort
}

func New(port ports.UserPort) *Handler {
	return &Handler{
		port,
	}
}

func (h *Handler) Create(writer http.ResponseWriter, request *http.Request) {
  ctx := context.Background()

  email := request.FormValue("email")

  user, err := h.port.Get(ctx, email)
  if err != nil {
    web.HandleError(err, writer)
    return
  }

  bytes, err := json.Marshal(user)
  if err != nil {
    web.HandleError(err, writer)
    return
  }

  writer.WriteHeader(http.StatusCreated)
  io.WriteString(writer, string(bytes))
}

func (h *Handler) Get(writer http.ResponseWriter, request *http.Request) {
  panic("not implemented")
}
