package config

import (
	"os"
	"strconv"
	"time"
)

const (
	HeaderRequestID = "X-Request-ID"
	RequestTimeout  = 60 * time.Second

	// CORS configuration
	CORSAllowedOrigins = "*"
	CORSAllowedMethods = "GET, POST, PUT, PATCH, DELETE, OPTIONS"
	CORSAllowedHeaders = "Accept, Content-Type, Authorization, X-Request-ID"
	CORSMaxAge         = 86400 // 24 hours in seconds
)

type HttpServerConfig struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Environment  string
}

func Get() HttpServerConfig {
	return HttpServerConfig{
		Host:         getEnv("SERVER_HOST", "localhost"),
		Port:         getEnv("SERVER_PORT", "3000"),
		ReadTimeout:  getEnv("SERVER_READ_TIMEOUT", 10*time.Second),
		WriteTimeout: getEnv("SERVER_WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:  getEnv("SERVER_IDLE_TIMEOUT", 60*time.Second),
		Environment:  getEnv("APP_ENV", "development"),
	}
}

// getEnv retrieves an environment variable or returns a default value
// Supports string and time.Duration types
func getEnv[T any](key string, defaultValue T) T {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	var result any
	switch any(defaultValue).(type) {
	case string:
		result = value
	case time.Duration:
		// For Duration, expect value in seconds as integer
		if seconds, err := strconv.Atoi(value); err == nil {
			result = time.Duration(seconds) * time.Second
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}

	return result.(T)
}
