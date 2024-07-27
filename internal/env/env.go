package env

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func init() {
	slog.Info("Loading env variables")
	err := godotenv.Load()
	if err != nil {
		slog.Error("Failed to load .env", "error", err)
	}
}
