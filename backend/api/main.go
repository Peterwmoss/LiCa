package main

import (
	"context"
	"os"

	"github.com/Peterwmoss/LiCa/api"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
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

  err := api.Start(3000, ctx)
	if err != nil {
		log.Fatal().Msg("Failed to start api")
	}
}
