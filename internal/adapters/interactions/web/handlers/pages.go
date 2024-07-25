package handlers

import (
	"net/http"
)

func HtmlIndex() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		templates.Render(writer, "index", nil)
	})
}

