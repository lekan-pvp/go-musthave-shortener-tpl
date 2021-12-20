package controllers

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func (controller *Controller) GetURLByID(w http.ResponseWriter, r *http.Request) {
	log.Println("IN GetURLByID")
	short := chi.URLParam(r, "articleID")
	if short == "" {
		http.Error(w, "url is empty", 404)
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

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
