package custom_logger

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func CustomLogger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logger.Info("CustomLogger middleware initialized")
		fn := func(w http.ResponseWriter, r *http.Request) {
			init := logger.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("ip", r.RemoteAddr),
			)
			middle := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			defer func() {
				init.Info("Completed request",
					slog.Int("status", middle.Status()),
					slog.String("latency", time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(middle, r)
		}

		return http.HandlerFunc(fn)
	}
}
