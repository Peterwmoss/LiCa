package handlers

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/responses"
	"github.com/Peterwmoss/LiCa/internal/core"
)

func IsHTMXRequest(request *http.Request) bool {
	return request.Header.Get("HX-Request") != ""
}

func RunWithContext(writer http.ResponseWriter, request *http.Request, worker func(chan responses.Response)) {
	resultCh := make(chan responses.Response)

	go worker(resultCh)

	for {
		select {
		case <-request.Context().Done():
			slog.Warn("context canceled")
			return
		case res := <-resultCh:
			if res.Err == nil {
				writer.WriteHeader(res.StatusCode)
				io.WriteString(writer, res.Data)
				return
			}

			if errors.Is(res.Err, core.ErrValidation) {
				err := errors.Join(res.Err, ErrBadRequest)
				slog.Error("error", "message", res.Err)
				HandleError(writer, err)
				return
			}
      
			if errors.Is(res.Err, core.ErrNotFound) {
				err := errors.Join(res.Err, ErrNotFound)
				slog.Error("error", "message", res.Err)
				HandleError(writer, err)
				return
			}

			err := errors.Join(res.Err, ErrInternalServerError)
			slog.Error("error", "message", res.Err)
			HandleError(writer, err)
		}
	}
}
