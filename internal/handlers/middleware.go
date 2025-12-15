package handlers

import (
	"net/http"
	"time"

	"github.com/your-org/go-rest-layered-template/internal/logger"

	"github.com/go-chi/chi/v5/middleware"
)

func RequestLogger(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()

			next.ServeHTTP(ww, r)

			log.Info("request",
				logger.String("method", r.Method),
				logger.String("path", r.URL.Path),
				logger.Int("status", ww.Status()),
				logger.Int("bytes", ww.BytesWritten()),
				logger.String("request_id", middleware.GetReqID(r.Context())),
				logger.String("remote", r.RemoteAddr),
				logger.String("duration", time.Since(start).String()),
			)
		}
		return http.HandlerFunc(fn)
	}
}
