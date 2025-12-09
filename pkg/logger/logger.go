package logger

import (
	"log/slog"
	"os"
)

func NewLogger(env string) *slog.Logger {
	var level slog.Level
	var handler slog.Handler

	switch env {
	case "production":
		level = slog.LevelInfo
	case "development":
		level = slog.LevelDebug
	default:
		level = slog.LevelInfo
	}

	handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})

	return slog.New(handler)
}
