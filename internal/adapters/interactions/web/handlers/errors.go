package handlers

import (
	"errors"
	"io"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/core"
)

func HandleError(w http.ResponseWriter, err error) {
	if errors.Is(err, core.ErrValidation) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad Request")
		return
	}

	if errors.Is(err, core.ErrNotFound) || errors.Is(err, core.ErrAccessDenied) {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Either you don't have access or the thing you're looking for doesn't exist")
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "Internal Server Error")
}
