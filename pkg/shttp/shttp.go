package shttp

import (
	"log/slog"
	"ya-storage/internal/config"
	custom_logger "ya-storage/internal/transport/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Init(cfg *config.Config, logger *slog.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(custom_logger.CustomLogger(logger))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	return router
}
