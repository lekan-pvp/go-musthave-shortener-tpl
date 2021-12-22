package controllers

import (
	"context"
	"net/http"
	"time"
	_ "github.com/lib/pq"
)


func (controller *Controller) PingDBHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	err := controller.PingDB(ctx)

	if err != nil {
		if err = controller.CloseDB(); err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
}
