package config

import "time"

const HeaderRequestID = "X-Request-ID"

type HttpServerConfig struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func Get() HttpServerConfig {
	return HttpServerConfig{
		Host:         "localhost",
		Port:         "3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
