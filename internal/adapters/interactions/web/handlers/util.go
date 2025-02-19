package handlers

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/responses"
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

			slog.Error("error", "message", res.Err)
			HandleError(writer, res.Err)
		}
	}
}
