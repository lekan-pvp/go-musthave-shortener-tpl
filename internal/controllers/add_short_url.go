package controllers

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"github.com/go-musthave-shortener-tpl/internal/key_gen"
	"io"
	"net/http"
	"strings"
	"time"
)

func (controller *Controller) AddURL(w http.ResponseWriter, r *http.Request) {
	var uuid string
	var cookie *http.Cookie
	var err error

	cookie, err = r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie){
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

	ctx, stop := context.WithTimeout(r.Context(), 1*time.Second)
	defer stop()

	orig := string(body)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	short := key_gen.GenerateShortLink(orig, uuid)
	short, err = controller.InsertUser(ctx, uuid, short, orig)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(controller.Cfg.BaseURL + "/" + short))
}
