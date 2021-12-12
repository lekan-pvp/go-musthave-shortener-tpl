package controllers

import (
	"io"
	"net/http"
)

func (controller *URLsController) AddURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	url := string(body)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	key, err := controller.CreateURL(url)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(controller.Cfg.BaseURL + "/" + key))
}


