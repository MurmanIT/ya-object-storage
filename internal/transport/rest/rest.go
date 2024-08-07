package rest

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"ya-storage/internal/config"
	"ya-storage/internal/transport/rest/routers/upload"
	"ya-storage/pkg/shttp"

	"github.com/braintree/manners"
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

func (s *ServerRest) Run() {

	upload.UploadFiles(&s.cfg.S3, s.logger, s.router)

	port := fmt.Sprint(":", s.cfg.HttpServer.Port)
	s.logger.Info("Starting server", slog.String("port", port))
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go s.listenForShutdown(ch)
	manners.ListenAndServe(port, s.router)
}

func (s *ServerRest) listenForShutdown(ch <-chan os.Signal) {
	<-ch
	fmt.Println("\rshutting down server...")
	manners.Close()
}
