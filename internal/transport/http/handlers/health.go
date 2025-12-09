package handlers

import (
	"net/http"

	"github.com/shamssahal/go-server/pkg/utils"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthCheck is a simple liveness probe
// Returns 200 if the service process is running
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJson(w, http.StatusOK, HealthResponse{
		Status: "ok",
	})
}
