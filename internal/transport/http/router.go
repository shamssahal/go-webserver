package http

import (
	"net/http"

	"github.com/shamssahal/go-server/internal/transport/http/handlers"
	"github.com/shamssahal/go-server/internal/transport/http/middleware"
)

func Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	// Health checks (no middleware needed - must be fast)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/ready", handlers.ReadinessCheck)

	// Application routes
	mux.HandleFunc("/do", handlers.HandleDo)

	// Apply global middleware
	app := Chain(
		mux,
		middleware.RequestID,
		middleware.Recover,
		middleware.RequestLog,
	)
	return app
}
