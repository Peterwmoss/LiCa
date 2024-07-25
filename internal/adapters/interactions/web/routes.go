package web

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/auth"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/middleware"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type Router struct {
	UserService ports.UserService
}

func (r *Router) SetupRoutes(server *http.ServeMux) {
	oauthConfig := auth.NewOauth2Config("/auth")
	server.Handle("GET /auth/login", handlers.AuthLogin(oauthConfig))
	server.Handle("GET /auth/logout", handlers.AuthLogout())
	server.Handle("GET /auth/callback", handlers.AuthCallback(oauthConfig, r.UserService))

	authMiddleware := middleware.AuthMiddleware(r.UserService, "/auth/login")

	server.Handle("GET /public/", handlers.StaticHandler())
	server.Handle("GET /", authMiddleware(handlers.HtmlIndex()))

	//
	// server.Handle("GET /actions/lists/new", authMiddleware(handlers.NewList()))
	// server.Handle("GET /actions/lists/{id}/items", authMiddleware(handlers.NewItem()))
	// server.Handle("GET /actions/categories/options", authMiddleware(handlers.SelectCategory(categoryService)))
	//
	// server.Handle("GET /lists", authMiddleware(handlers.ListGetAll(listService)))
	// server.Handle("GET /lists/{id}", authMiddleware(handlers.ListGet(listService)))
	// server.Handle("POST /lists", authMiddleware(handlers.ListCreate(listService)))
	//
	// server.Handle("POST /lists/{id}/items", authMiddleware(handlers.ListItemCreate(productService, categoryService, listService, listItemService)))
	//
	// server.Handle(authGetUrl+"login", handlers.AuthLogin(authConfig, authStateCheck))
	// server.Handle(authGetUrl+"logout", handlers.AuthLogout())
	// server.Handle(authGetUrl+"callback", handlers.AuthCallback(authConfig, authStateCheck, userService))
}
