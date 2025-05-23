package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func ZapLogger(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := &statusWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}
			next.ServeHTTP(ww, r)

			duration := time.Since(start)

			logger.Info("HTTP request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status", ww.status),
				zap.Duration("duration", duration),
				zap.String("remote_addr", r.RemoteAddr),
			)
		})
	}
}
