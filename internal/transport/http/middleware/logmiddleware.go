package middleware

import (
	"log"
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
		// Very lightweight, structured-enough log line.
		log.Printf(
			`request_id=%s method=%s path=%q status=%d bytes=%d dur=%s remote=%s ua=%q`,
			rid,
			r.Method,
			r.URL.Path,
			rec.status,
			rec.size,
			latency,
			r.RemoteAddr,
			r.UserAgent(),
		)
	})
}
