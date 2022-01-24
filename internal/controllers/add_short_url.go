package controllers

import (
	"context"
	"github.com/jackc/pgerrcode"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/cookieserver"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/keygen"
	"github.com/lib/pq"
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

	ctx, stop := context.WithTimeout(r.Context(), 1*time.Second)
	defer stop()

	orig := string(body)

	status := http.StatusCreated
	short := keygen.GenerateShortLink(orig, uuid)
	short, err = controller.InsertUser(ctx, uuid, short, orig)
	if err != nil {
		if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
			status = http.StatusConflict
		} else {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(controller.Cfg.BaseURL + "/" + short))
}
