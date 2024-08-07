package s3

import (
	"log/slog"
	"sync"
	"ya-storage/internal/config"
)

type S3Handler struct {
	S3     *config.S3
	Logger *slog.Logger
}

var wg sync.WaitGroup

func Init(ConfigS3 *config.S3, logger *slog.Logger) *S3Handler {
	return &S3Handler{
		S3:     ConfigS3,
		Logger: logger,
	}
}
