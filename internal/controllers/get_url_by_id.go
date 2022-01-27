package controllers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
	"log"
	"net/http"
)

func (service *Controller) GetURLByID(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "articleID")
	if short == "" {
		http.Error(w, "url is empty", 404)
		return
	}

	log.Println(short)

	orig := &models.OriginLink{}

	orig, err := service.GetOrigByShort(r.Context(), short)
	if err != nil {
		log.Println("IN ERR HANDLER GetOrigByShort")
		http.Error(w, err.Error(), 404)
		return
	}

	if orig == nil {
		http.NotFound(w, r)
		return
	}

	if !orig.IsDeleted() {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", orig.Link)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if orig.IsDeleted() {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(410)
		return
	}
}
