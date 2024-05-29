package web

import (
	"net/http"

)

func setupRoutes(server *http.ServeMux) {
	server.Handle("GET /public/", staticHandler)

	server.Handle("GET /", authMiddleware(handlers.GetIndex()))

	server.Handle("GET /actions/lists/new", authMiddleware(handlers.NewList()))
	server.Handle("GET /actions/lists/{id}/items", authMiddleware(handlers.NewItem()))
	server.Handle("GET /actions/categories/options", authMiddleware(handlers.SelectCategory(categoryService)))

	server.Handle("GET /lists", authMiddleware(handlers.ListGetAll(listService)))
	server.Handle("GET /lists/{id}", authMiddleware(handlers.ListGet(listService)))
	server.Handle("POST /lists", authMiddleware(handlers.ListCreate(listService)))

	server.Handle("POST /lists/{id}/items", authMiddleware(handlers.ListItemCreate(productService, categoryService, listService, listItemService)))

	server.Handle(authGetUrl+"login", handlers.AuthLogin(authConfig, authStateCheck))
	server.Handle(authGetUrl+"logout", handlers.AuthLogout())
	server.Handle(authGetUrl+"callback", handlers.AuthCallback(authConfig, authStateCheck, userService))
}
