package main

import (
	"fmt"
	"log/slog"
	"ya-storage/internal/config"
	"ya-storage/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Errorf("config error: %s", err)
		return
	}
	logger := logger.Init(cfg)
	logger = logger.With(
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.HttpServer.Port),
	)
	logger.Info("Logger initialized")
}

// TODO: init server
// TODO: init storage
// TODO: run server
