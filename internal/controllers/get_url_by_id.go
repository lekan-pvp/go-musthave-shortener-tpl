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

	orig, err := controller.GetOrigByShort(r.Context(), short)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	if orig == "" {
		http.NotFound(w, r)
		return
	}

	log.Printf("In get_url_by_id: %s\n", orig)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Location", orig)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
