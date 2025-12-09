package main

import (
	"github.com/shamssahal/go-server/config"
)

type APIServer struct {
	cfg config.HttpServerConfig
}

func NewAPIServer(cfg config.HttpServerConfig) *APIServer {
	return &APIServer{
		cfg: cfg,
	}
}
