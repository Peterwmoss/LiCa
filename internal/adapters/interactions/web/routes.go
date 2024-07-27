package web

import (
	"net/http"

	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/auth"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers/actions"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers/components"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/handlers/pages"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/middleware"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type Router struct {
	UserService     ports.UserService
	ListService     ports.ListService
	ListItemService ports.ListItemService
	ProductService  ports.ProductService
	CategoryService ports.CategoryService
}

func (r *Router) SetupRoutes(server *http.ServeMux) {
	oauthConfig := auth.NewOauth2Config("/auth")
	server.Handle("GET /auth/login", handlers.AuthLogin(oauthConfig))
	server.Handle("GET /auth/logout", handlers.AuthLogout())
	server.Handle("GET /auth/callback", handlers.AuthCallback(oauthConfig, r.UserService))

	authMiddleware := middleware.AuthMiddleware(r.UserService, "/auth/login", oauthConfig)

	indexPage := pages.Index{}
	listPage := pages.List{ListService: r.ListService}

	listAction := actions.List{ListService: r.ListService, ListPage: listPage}
	listItemAction := actions.ListItem{ListItemService: r.ListItemService, ListPage: listPage}

	listComponent := components.List{}
	listItemsComponent := components.ListItem{CategoryService: r.CategoryService}

	server.Handle("GET /public/", handlers.StaticHandler())

	server.Handle("GET /", authMiddleware(http.HandlerFunc(indexPage.Index)))

	server.Handle("GET /components/lists/new", authMiddleware(http.HandlerFunc(listComponent.New)))
	server.Handle("GET /components/items/new", authMiddleware(http.HandlerFunc(listItemsComponent.New)))

	server.Handle("POST /actions/lists", authMiddleware(http.HandlerFunc(listAction.Create)))
	server.Handle("POST /actions/items", authMiddleware(http.HandlerFunc(listItemAction.Add)))

	server.Handle("GET /pages/lists", authMiddleware(http.HandlerFunc(listPage.Lists)))
	server.Handle("GET /pages/lists/{name}", authMiddleware(http.HandlerFunc(listPage.List)))
}
