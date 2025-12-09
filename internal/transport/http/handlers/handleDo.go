package handlers

import (
	"log"
	"net/http"

	"github.com/shamssahal/go-server/pkg/utils"
)

func HandleDo(w http.ResponseWriter, r *http.Request) {
	rid := utils.RequestIDFromContext(r.Context())
	log.Printf("handling /do, request_id=%s", rid)
	utils.WriteJson(w, http.StatusOK, struct{}{})

}
