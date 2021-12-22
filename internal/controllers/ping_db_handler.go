package controllers

import (
	"context"
	"net/http"
	"time"
)


func (controller *Controller) PingDBHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	err := controller.PingDB(ctx)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
}
