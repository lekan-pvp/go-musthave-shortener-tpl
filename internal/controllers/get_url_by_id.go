package controllers

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func (controller *Controller) GetURLByID(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "articleID")
	if short == "" {
		http.Error(w, "url is empty", 404)
		return
	}

	log.Println(short)

	orig, err := controller.GetOrigByShort(r.Context(), short)
	if err != nil {
		log.Println("IN ERR HANDLER GetOrigByShort")
		http.Error(w, err.Error(), 404)
		return
	}

	if orig == "" {
		http.NotFound(w, r)
		return
	}

	if orig != "deleted" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", orig)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}


	if orig == "deleted" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(410)
		if err = controller.DeleteItem(r.Context(), short); err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		return
	}
}
