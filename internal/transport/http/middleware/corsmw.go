package middleware

import (
	"net/http"
	"strconv"

	"github.com/shamssahal/go-server/config"
)

// CORS middleware handles Cross-Origin Resource Sharing
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", config.CORSAllowedOrigins)

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", config.CORSAllowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", config.CORSAllowedHeaders)
			w.Header().Set("Access-Control-Max-Age", strconv.Itoa(config.CORSMaxAge))
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}
