package controllers

import (
	"context"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func (service *Controller) PingDBHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 250*time.Millisecond)
	defer cancel()

	err := service.CheckPing(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
}
