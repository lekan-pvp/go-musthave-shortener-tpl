package controllers

import (
	"encoding/json"
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"github.com/go-musthave-shortener-tpl/internal/models"
	"log"
	"net/http"
	"strings"
)


var out []models.URLs

func (controller *URLsController) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uid")
	if err != nil {
		cookie = cookie_handler.CreateCookie()
		http.SetCookie(w, cookie)
	}

	if cookie_handler.CheckCookie(cookie) {
		cookie = cookie_handler.CreateCookie()
		http.SetCookie(w, cookie)
	}

	values := strings.Split(cookie.Value, ":")
	uuid := values[0]

	out = controller.GetURLsListByUUID(uuid, controller.Cfg.BaseURL)



	log.Println(out)

	result, err := json.MarshalIndent(&out, "", " ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if len(out) == 0 {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(200)
	}

	w.Write(result)
}
