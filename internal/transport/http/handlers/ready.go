package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/shamssahal/go-server/pkg/utils"
)

// ReadinessResponse represents the readiness check response
type ReadinessResponse struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks,omitempty"`
}

// ReadinessCheck checks if the service is ready to handle traffic
// Checks dependencies like database, external services, etc.
func ReadinessCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	checks := make(map[string]string)
	allHealthy := true

	// if err := checkDatabase(ctx); err != nil {
	//     checks["database"] = "unhealthy: " + err.Error()
	//     allHealthy = false
	// } else {
	//     checks["database"] = "healthy"
	// }

	checks["service"] = "healthy"

	status := http.StatusOK
	statusText := "ready"
	if !allHealthy {
		status = http.StatusServiceUnavailable
		statusText = "not ready"
		slog.WarnContext(ctx, "readiness check failed", "checks", checks)
	}

	utils.WriteJson(w, status, ReadinessResponse{
		Status: statusText,
		Checks: checks,
	})
}
