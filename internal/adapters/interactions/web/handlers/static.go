package handlers

import "net/http"

func StaticHandler() http.Handler {
	return http.FileServer(http.Dir("./internal/adapters/interactions/web/assets"))
}
