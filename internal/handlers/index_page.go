package handlers

import (
	"net/http"
)

func GetIndex() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		templates.Render(writer, "index", nil)
	})
}
