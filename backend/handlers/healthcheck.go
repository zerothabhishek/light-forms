package handlers

import (
	"net/http"

	"github.com/light-forms/backend/services"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthMsg, err := services.HealthCheck()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(healthMsg))
}
