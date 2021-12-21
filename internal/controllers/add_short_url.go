package controllers

import (
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"io"
	"log"
	"net/http"
	"strings"
)

func (controller *Controller) AddURL(w http.ResponseWriter, r *http.Request) {
	var uuid string
	var cookie *http.Cookie
	var err error

	cookies := r.Cookies()
	for c := range cookies {
		log.Println(c)
	}

	cookie, err = r.Cookie("token")
	if err != nil {
		cookie = cookie_handler.CreateCookie()
		http.SetCookie(w, cookie)
	}

	if !cookie_handler.CheckCookie(cookie) {
		cookie = cookie_handler.CreateCookie()
		http.SetCookie(w, cookie)
	}

	values := strings.Split(cookie.Value, ":")
	uuid = values[0]

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
