package handlers

import (
	"errors"
	"io"
	"net/http"
)

var (
	ErrBadRequest          = errors.New("Bad Request")
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Not Found")
	ErrForbidden           = errors.New("Forbidden")
)

func HandleError(w http.ResponseWriter, err error) {
	if errors.Is(err, ErrBadRequest) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad Request")
		return
	}

	if errors.Is(err, ErrNotFound) || errors.Is(err, ErrForbidden) {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Either you don't have access or the thing you're looking for doesn't exist")
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "Internal Server Error")
}
