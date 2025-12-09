package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

// Use an unexported key type to avoid collisions in context.
type ctxKey struct{}

var requestIDKey = &ctxKey{}

func NewRequestID() string {
	var b [16]byte
	_, _ = rand.Read(b[:]) // best effort; if it fails, ID will be zero bytes which is fine for a fallback
	return hex.EncodeToString(b[:])
}

func RequestIDFromContext(ctx context.Context) string {
	if v := ctx.Value(requestIDKey); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// WithRequestID returns a new Context carrying the request ID.
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
