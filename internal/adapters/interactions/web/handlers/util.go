package handlers

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/responses"
)

func isHTMXRequest(request *http.Request) bool {
	return request.Header.Get("HX-Request") != ""
}

func runWithContext(writer http.ResponseWriter, request *http.Request, worker func(chan responses.Response)) {
	resultCh := make(chan responses.Response)

  go worker(resultCh)

	for {
		select {
		case <-request.Context().Done():
			slog.Warn("context canceled")
			return
		case res := <-resultCh:
			if res.Err() != nil {
        slog.Error("error", "error", res.Err())
        writer.WriteHeader(500)
        io.WriteString(writer, "Internal server error")
				return
			}

			writer.WriteHeader(res.StatusCode())
			io.WriteString(writer, res.Message())
		}
	}

}
