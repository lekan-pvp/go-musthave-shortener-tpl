package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-musthave-shortener-tpl/internal/config"
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
	"net/http"
)

type URLsController struct {
	interfaces.IURLsService
	Cfg *config.Config
}

func (controller *URLsController) GetURLByID(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "articleID")
	if key == "" {
		http.Error(w, "url not found", 404)
		return
	}
	url, err := controller.GetURLsDetail(key)
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
