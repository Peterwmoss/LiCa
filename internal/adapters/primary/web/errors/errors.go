package web

import (
	"errors"
	"io"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal"
)

func HandleError(err error, writer http.ResponseWriter) {
	switch {
	case errors.Is(err, core.ErrBadRequest):
		writeError(writer, http.StatusBadRequest, errors.Unwrap(err).Error())
	}

	writeError(writer, http.StatusInternalServerError, "internal server error")
}

func writeError(writer http.ResponseWriter, status int, message string) {
  writer.WriteHeader(status)
  io.WriteString(writer, message)
}
