package controllers

import (
	"encoding/json"
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"github.com/go-musthave-shortener-tpl/internal/models"
	"io"
	"net/http"
)

func (controller *Controller) ApiShortenBatch(w http.ResponseWriter, r *http.Request) {
	in := make([]models.BatchIn, 0)

	cookie, err := r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie) {
		cookie = cookie_handler.CreateCookie()
	}

	http.SetCookie(w, cookie)

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(body, &in)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	result := controller.BanchApi(r.Context(), in, controller.Cfg.BaseURL)

	inBody, err := json.Marshal(&result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	w.Write(inBody)
}
