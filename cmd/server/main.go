package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/Peterwmoss/LiCa/internal/domain"
	"github.com/Peterwmoss/LiCa/internal/functions/auth"
	"github.com/Peterwmoss/LiCa/internal/functions/home"
	"github.com/Peterwmoss/LiCa/internal/functions/list"
	"github.com/Peterwmoss/LiCa/internal/functions/user"
	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()
	db := database.Get()
	defer db.Close()

  if err := database.CreateSchema(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to create schema")
  }

  if err := database.Seed(db, ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to seed data")
  }

  app.Static("/public", "./internal/assets/public")

  homeHandler := home.NewHandler(domain.NewUserService())
  homeHandler.Mount(app)

  listHandler := list.NewHandler(db, ctx)
  listHandler.Mount(app)

  authHandler := auth.NewHandler(db, ctx)
  authHandler.Mount(app)

  userHandler := user.NewHandler()
  userHandler.Mount(app)

  if err := app.Listen(":3000"); err != nil {
		log.Fatal().Msg("Failed to start api")
	}
}
