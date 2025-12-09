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
	root := http.NewServeMux()

	// health mux with no middleware overheads
	healthMux := http.NewServeMux()
	healthMux.HandleFunc("GET /health", handlers.HealthCheck)
	healthMux.HandleFunc("GET /ready", handlers.ReadinessCheck)

	// app mux with routes & middlewares
	appMux := http.NewServeMux()
	appMux.HandleFunc("GET /do", handlers.HandleDo)
	app := Chain(
		appMux,
		middleware.CORS,
		middleware.RequestTimeout,
		middleware.RequestID,
		middleware.Recover,
		middleware.RequestLog,
	)

	root.Handle("/health", healthMux)
	root.Handle("/ready", healthMux)
	root.Handle("/", app)

	return root
}
