package controllers

import (
	"encoding/json"
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"github.com/go-musthave-shortener-tpl/internal/models"
	"io"
	"net/http"
	"strings"
)

func (controller *Controller) ApiShortenBatch(w http.ResponseWriter, r *http.Request) {
	var uuid string

	in := make([]models.BatchIn, 0)

	cookie, err := r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie) {
		cookie = cookie_handler.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	uuid = values[0]

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

	result, err := controller.BanchApi(r.Context(), uuid, in, controller.Cfg.BaseURL)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	inBody, err := json.Marshal(&result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	w.Write(inBody)
}
