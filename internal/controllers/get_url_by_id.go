package controllers

import (
	"context"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

func (controller *Controller) GetURLByID(w http.ResponseWriter, r *http.Request) {
	log.Println("IN GetURLByID")
	short := chi.URLParam(r, "articleID")
	if short == "" {
		http.Error(w, "url is empty", 404)
		return
	}

	ctx, stop := context.WithTimeout(r.Context(), 1*time.Second)
	defer stop()

	orig, err := controller.GetOrigByShortDB(ctx, short)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	if orig == "" {
		http.NotFound(w, r)
		return
	}

	url, err := controller.GetURLs(short)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if url == "" {
		http.NotFound(w, r)
		return
	}

	log.Printf("In get_url_by_id: %s == %s\n", orig, url)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Location", orig)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
