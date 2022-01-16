package controllers

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func (controller *Controller) GetURLByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("IN GetURLByID controller is %T", controller)
	short := chi.URLParam(r, "articleID")
	if short == "" {
		http.Error(w, "url is empty", 404)
		return
	}

	log.Println(short)

	orig, err := controller.GetOrigByShort(r.Context(), short)
	log.Println("ORIG:", orig)
	if err != nil {
		log.Println("IN ERR HANDLER GetOrigByShort")
		http.Error(w, err.Error(), 404)
		return
	}

	if orig == "" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Location", orig)
	w.WriteHeader(http.StatusTemporaryRedirect)

	if orig == "deleted" {
		w.WriteHeader(410)
	}

	log.Printf("In get_url_by_id: %s\n", orig)
}
