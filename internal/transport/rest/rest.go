package rest

import (
	"log/slog"
	"ya-storage/internal/config"
	"ya-storage/pkg/shttp"

	"github.com/go-chi/chi/v5"
)

type ServerRest struct {
	cfg    *config.Config
	router *chi.Mux
	logger *slog.Logger
}

var server ServerRest

func Init(cfg *config.Config, logger *slog.Logger) *ServerRest {
	server.cfg = cfg
	server.logger = logger
	server.router = shttp.Init(cfg, logger)

	return &server
}
