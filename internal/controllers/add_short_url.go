package controllers

import (
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"io"
	"net/http"
	"strings"
)

func (controller *URLsController) AddURL(w http.ResponseWriter, r *http.Request) {
	var uuid string

	cookie, err := r.Cookie("uid")
	if err != nil || !cookie_handler.CheckCookie(cookie){
		cookie = cookie_handler.CreateCookie()
		http.SetCookie(w, cookie)
	}


	values := strings.Split(cookie.Value, ":")
	uuid = values[0]


	//uuid = "123456789"
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	orig := string(body)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	short, err := controller.CreateURL(uuid, orig)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(controller.Cfg.BaseURL + "/" + short))
}
