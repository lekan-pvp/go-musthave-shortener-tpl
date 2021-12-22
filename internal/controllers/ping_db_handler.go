package controllers

import (
	"context"
	"net/http"
	"time"
)



func (controller *Controller) PingDBHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := controller.PingDB(ctx); err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
