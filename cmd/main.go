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
	"github.com/Peterwmoss/LiCa/internal/middleware"
	"github.com/Peterwmoss/LiCa/internal/views"

	"github.com/Peterwmoss/LiCa/internal/functions"
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

	templates := views.NewTemplates()

	userService := domain.NewUserService(db, ctx)
	categoryService := domain.NewCategoryService(db, ctx)
	productService := domain.NewProductService(db, ctx, categoryService)
	listItemService := domain.NewListItemService(productService)
	listService := domain.NewListService(db, ctx, listItemService)

	fileServer := http.FileServer(http.Dir("./internal/assets"))
	server.Handle("GET /public/", fileServer)

	server.Handle("GET /", middleware.UseAuth(userService, functions.GetIndex(templates)))

	listHandler := functions.NewListHandler(listService, templates)
	server.Handle("GET /lists", middleware.UseAuth(userService, listHandler.GetAll))
  server.Handle("GET /lists/{id}", middleware.UseAuth(userService, listHandler.Get))
	server.Handle("POST /lists", middleware.UseAuth(userService, listHandler.Create))

	userHandler := functions.NewUserHandler(userService, templates)
	server.Handle("GET /users", userHandler.Get)

	authConfig := auth.NewAuthConfig("/auth")
	authHandler := functions.NewAuthHandler(userService, authConfig.BaseUrl)

	server.Handle("GET /auth/login", authHandler.Login)
	server.Handle("GET /auth/logout", authHandler.Logout)
	server.Handle("GET /auth/callback", authHandler.Callback)

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatal().Msg("Failed to start api")
	}
}
