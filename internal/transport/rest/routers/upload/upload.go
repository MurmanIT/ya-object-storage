package upload

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"ya-storage/internal/transport/rest/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func UploadFiles(log *slog.Logger, router *chi.Mux) {
	log = log.With(
		slog.String("component", "upload"),
		slog.String("method", "UploadFiles"),
	)
	router.Route("/upload", func(r chi.Router) {
		r.Put("/", uploadFile(log))
	})
}

func uploadFile(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filepath, err := fileSendToServer(r, log)
		if err != nil {
			render.JSON(w, r, response.Error(
				"Failed to upload file "+err.Error(),
			))
			return
		}
		log.Info("File uploaded", slog.String("path", filepath))
	}
}

func fileSendToServer(r *http.Request, log *slog.Logger) (string, error) {
	log.Info("Intialized uploadFile")
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Error("Failed to get file", slog.String("error", err.Error()))
		return "", err
	}
	defer file.Close()
	log.Info("Upload file", slog.String("name", handler.Filename))
	log.Info("File size", slog.Int64("size", handler.Size))

	f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error("Failed to open file", slog.String("error", err.Error()))
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	filePath := "./files/" + handler.Filename
	return filePath, nil
}
