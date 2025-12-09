package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/shamssahal/go-server/config"
	httpserver "github.com/shamssahal/go-server/internal/transport/http"
	"github.com/shamssahal/go-server/pkg/logger"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Get()

	s := NewAPIServer(cfg)
	logger := logger.NewLogger(cfg.Environment)
	slog.SetDefault(logger)

	srv := &http.Server{
		Addr:         net.JoinHostPort(s.cfg.Host, s.cfg.Port),
		Handler:      httpserver.NewHandler(),
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
		IdleTimeout:  s.cfg.IdleTimeout,
	}

	errCh := make(chan error, 1)
	go func() {
		slog.Info("server starting", "addr", srv.Addr, "env", cfg.Environment)
		errCh <- srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	select {
	case sig := <-stop:
		slog.Info("shutdown signal received", "signal", sig.String())
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			slog.Error("server failed to start", "error", err)
			os.Exit(1)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("graceful shutdown failed", "error", err, "timeout", "10s")
		os.Exit(1)
	}
	slog.Info("server stopped gracefully")

}
