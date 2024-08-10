package upload

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"ya-storage/internal/config"
	"ya-storage/internal/transport/rest/response"
	"ya-storage/pkg/s3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func UploadFiles(ConfigS3 *config.S3, log *slog.Logger, router *chi.Mux) {
	log = log.With(
		slog.String("component", "upload"),
		slog.String("method", "UploadFiles"),
	)
	router.Route("/upload", func(r chi.Router) {
		r.Put("/", uploadFile(log, ConfigS3))
		r.Get("/", printObject(log, ConfigS3))
		r.Delete("/", deleteObject(log, ConfigS3))
	})
}

func printObject(log *slog.Logger, ConfigS3 *config.S3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func deleteObject(log *slog.Logger, ConfigS3 *config.S3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func uploadFile(log *slog.Logger, ConfigS3 *config.S3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filepath, name, err := fileSendToServer(r, log)

		if filepath == "" && name == "" && err != nil {
			render.JSON(w, r, response.Error(
				"Failed to upload file "+err.Error(),
			))
			return
		}

		sh3 := s3.Init(ConfigS3, log)
		sh3.UploadFile(filepath, name)

		if err != nil {
			render.JSON(w, r, response.Error(
				"Failed to upload file "+err.Error(),
			))
			return
		}
		log.Info("File uploaded", slog.String("path", filepath))
	}
}

func fileSendToServer(r *http.Request, log *slog.Logger) (filePath string, name string, err error) {
	log.Info("Intialized uploadFile")
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Error("Failed to get file", slog.String("error", err.Error()))
		return "", "", err
	}
	defer file.Close()
	log.Info("Upload file", slog.String("name", handler.Filename))
	log.Info("File size", slog.Int64("size", handler.Size))

	f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error("Failed to open file", slog.String("error", err.Error()))
		return "", "", err
	}
	defer f.Close()
	io.Copy(f, file)

	filePath = "./files/" + handler.Filename

	return filePath, handler.Filename, nil
}
