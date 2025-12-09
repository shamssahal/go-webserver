package middleware

import (
	"net/http"

	"github.com/shamssahal/go-server/config"
	"github.com/shamssahal/go-server/pkg/utils"
)

// RequestID extracts an incoming X-Request-ID if present.
// If missing, it generates one at the edge. It attaches the value to the
// request Context and also sets it on the response headers so clients can log it.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get(config.HeaderRequestID)
		if rid == "" {
			rid = utils.NewRequestID()
		}

		// Expose the correlation ID to downstream and back to the caller.
		ctx := utils.WithRequestID(r.Context(), rid)
		w.Header().Set(config.HeaderRequestID, rid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
