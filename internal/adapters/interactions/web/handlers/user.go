package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/responses"
	"github.com/Peterwmoss/LiCa/internal/core/domain"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) UserHandler {
	return UserHandler{
		userService,
	}
}

func (h *UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	runWithContext(writer, request, func(c chan responses.Response) {
    email := request.Context().Value("user").(domain.User).Email
		user, err := h.userService.Create(request.Context(), string(email))
		if err != nil {
			c <- responses.GenericResponse{Error: err}
			return
		}

		bytes, err := json.Marshal(user)
		if err != nil {
			c <- responses.GenericResponse{Error: err}
			return
		}

		c <- responses.GenericResponse{
			Msg:    string(bytes),
			Status: http.StatusCreated,
		}
	})
}

func (h *UserHandler) Get(writer http.ResponseWriter, request *http.Request) {
	runWithContext(writer, request, func(c chan responses.Response) {
    email := request.Context().Value("user").(domain.User).Email
		user, err := h.userService.Get(request.Context(), string(email))
		if err != nil {
			c <- responses.GenericResponse{Error: err}
			return
		}

		bytes, err := json.Marshal(user)
		if err != nil {
			c <- responses.GenericResponse{Error: err}
			return
		}

		c <- responses.GenericResponse{
			Msg:    string(bytes),
			Status: http.StatusOK,
		}
	})
}
