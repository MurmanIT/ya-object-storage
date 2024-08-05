package s3

import (
	"log/slog"
	"ya-storage/internal/config"
)

type S3Handler struct {
	//Session *session.Session
}

func Init(cfg *config.Config, logger *slog.Logger) *S3Handler {
	return &S3Handler{}
}
