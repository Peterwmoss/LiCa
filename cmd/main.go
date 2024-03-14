package main

import (
	"context"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/Peterwmoss/LiCa/internal/auth"
	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/functions"
	"github.com/Peterwmoss/LiCa/internal/middleware"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	log.Info().Msg("Loading env variables")
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Failed to load .env: %s", err)
	}
}

func main() {
	ctx := context.Background()
	db := database.Get()
	defer db.Close()

	if err := database.CreateSchema(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to create schema")
	}

	if err := database.Seed(db, ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to seed data")
	}

	server := http.NewServeMux()

	userService := domain.NewUserService(db, ctx)
	categoryService := domain.NewCategoryService(db, ctx)
	productService := domain.NewProductService(db, ctx, categoryService)
	listItemService := domain.NewListItemService(productService)
	listService := domain.NewListService(db, ctx, listItemService)

	fileServer := http.FileServer(http.Dir("./internal/assets/public"))
	server.Handle("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
    middleware.NewAuth(userService)(writer, request)

    fileServer.ServeHTTP(writer, request)
  }))

	homeHandler := functions.NewHomeHandler(userService)
	server.HandleFunc("GET /home", homeHandler.Get)

	listHandler := functions.NewListHandler(listService, userService)
	server.HandleFunc("GET /lists", listHandler.GetAll)
	server.HandleFunc("POST /lists", listHandler.Create)

	authConfig := auth.NewAuthConfig("/auth")
	authHandler := functions.NewAuthHandler(userService, authConfig.BaseUrl)

	server.HandleFunc("GET /auth/login", authHandler.Login)
	server.HandleFunc("GET /auth/logout", authHandler.Logout)
	server.HandleFunc("GET /auth/callback", authHandler.Callback)

	userHandler := functions.NewUserHandler(userService)
	server.HandleFunc("GET /users", userHandler.Get)

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatal().Msg("Failed to start api")
	}
}
