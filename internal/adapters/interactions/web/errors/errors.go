package web

import (
	"errors"
	"io"
	"net/http"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal error")
)

func HandleError(err error, writer http.ResponseWriter) {
	switch {
	case errors.Is(err, ErrBadRequest):
		writeError(writer, http.StatusBadRequest, errors.Unwrap(err).Error())
    return
	}

	writeError(writer, http.StatusInternalServerError, "internal server error")
}

func writeError(writer http.ResponseWriter, status int, message string) {
	writer.WriteHeader(status)
	io.WriteString(writer, message)
}
