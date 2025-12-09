package middleware

import (
	"log/slog"
	"net/http"

	"github.com/shamssahal/go-server/pkg/utils"
)

// Recover recovers from panics and returns a 500 Internal Server Error.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				slog.Error("panic recovered",
					"error", rec,
					slog.String("requestId", utils.RequestIDFromContext(r.Context())),
					slog.String("path", r.URL.Path),
					slog.String("method", r.Method),
				)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
