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
	"github.com/Peterwmoss/LiCa/internal/handlers"
	"github.com/Peterwmoss/LiCa/internal/handlers/middleware"
)


var authStateCheck string

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	log.Info().Msg("Loading env variables")
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Failed to load .env: %s", err)
	}

	logLevel, present := os.LookupEnv("LICA_LOG_LEVEL")
	if !present {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		level, err := zerolog.ParseLevel(logLevel)
		if err != nil {
			panic("invalid log level")
		}
		zerolog.SetGlobalLevel(level)
	}

	stateCheck, present := os.LookupEnv("LICA_STATE_CHECK")
	if !present {
		stateCheck = "a8e7hfwnkf3"
	}
	authStateCheck = stateCheck
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

	userService := domain.NewUserService(db, ctx)
	categoryService := domain.NewCategoryService(db, ctx)
	productService := domain.NewProductService(db, ctx, categoryService)
	listItemService := domain.NewListItemService(productService)
	listService := domain.NewListService(db, ctx, listItemService)

	staticHandler := http.FileServer(http.Dir("./internal/assets"))

  const authBaseUrl = "/auth"
  const authGetUrl = "GET " + authBaseUrl + "/"

	authConfig := auth.NewOauth2Config(authBaseUrl)
	authMiddleware := middleware.AuthMiddleware(userService, authBaseUrl+"/login")

	server := http.NewServeMux()

	server.Handle("GET /public/", staticHandler)

	server.Handle("GET /", authMiddleware(handlers.GetIndex()))

	server.Handle("GET /lists", authMiddleware(handlers.ListGetAll(listService)))
	server.Handle("GET /lists/{id}", authMiddleware(handlers.ListGet(listService)))
	server.Handle("POST /lists", authMiddleware(handlers.ListCreate(listService)))

	server.Handle(authGetUrl+"login", handlers.AuthLogin(authConfig, authStateCheck))
	server.Handle(authGetUrl+"logout", handlers.AuthLogout())
	server.Handle(authGetUrl+"callback", handlers.AuthCallback(authConfig, authStateCheck, userService))

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatal().Msg("Failed to start api")
	}
}
