package handlers

import (
	"net/http"

	"github.com/shamssahal/go-server/pkg/utils"
)

func HandleDo(w http.ResponseWriter, r *http.Request) {
	//do business logic here
	utils.WriteJson(w, http.StatusOK, struct{}{})

	//if err use
	// utils.WriteError(w, http.StatusBadRequest,struct{}{})

}
