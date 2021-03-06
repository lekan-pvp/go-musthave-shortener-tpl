package controllers

import (
	"encoding/json"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/cookieserver"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
	"io"
	"net/http"
	"strings"
)

func (service *Controller) APIShortenBatch(w http.ResponseWriter, r *http.Request) {
	var uuid string

	in := make([]models.BatchIn, 0)

	cookie, err := r.Cookie("token")
	if err != nil || !cookieserver.CheckCookie(cookie) {
		cookie = cookieserver.CreateCookie()
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

	result, err := service.BanchAPI(r.Context(), uuid, in, service.Cfg.BaseURL)
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
	w.WriteHeader(201)

	w.Write(inBody)
}
