package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/shamssahal/go-server/pkg/utils"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *statusRecorder) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusRecorder) Write(b []byte) (int, error) {
	if w.status == 0 {
		// If WriteHeader wasn't called, status is 200 by default.
		w.status = http.StatusOK
	}
	n, err := w.ResponseWriter.Write(b)
	w.size += n
	return n, err
}

// RequestLog logs request details including method, path, status, duration, etc.
func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rid := utils.RequestIDFromContext(r.Context())

		rec := &statusRecorder{ResponseWriter: w}
		next.ServeHTTP(rec, r)

		latency := time.Since(start)
		slog.LogAttrs(
			r.Context(),
			slog.LevelInfo,
			"[Request Log]",
			slog.String("requestId", rid),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rec.status),
			slog.Int("size", rec.size),
			slog.String("duration", latency.String()),
			slog.String("remote", r.RemoteAddr),
			slog.String("userAgent", r.UserAgent()),
		)
	})
}
