package logger

import (
	"log/slog"
	"os"
	"ya-storage/internal/config"
)

const (
	local = "local"
	dev   = "dev"
	prod  = "prod"
)

func Init(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case local:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case dev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case prod:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}
	return log
}
