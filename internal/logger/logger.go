package logger

import (
	"log/slog"
	"os"
	"strconv"
)

func init() {
	logLevel, present := os.LookupEnv("LICA_LOG_LEVEL")
	if !present {
		slog.SetLogLoggerLevel(slog.LevelDebug.Level())
	} else {
		parsed, err := strconv.Atoi(logLevel)
		if err != nil {
			slog.Error("Failed to parse log level", "error", err)
			return
		}

		slog.SetLogLoggerLevel(slog.Level(parsed))
	}

	slog.Info("logger initialized")
}
