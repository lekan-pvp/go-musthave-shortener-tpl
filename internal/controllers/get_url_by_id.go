package controllers

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/cookie_handler"
	repository_db "github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/repository-db"
	"log"
	"net/http"
	"strings"
)

func (controller *Controller) GetURLByID(w http.ResponseWriter, r *http.Request) {
	var uuid string

	cookie, err := r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie) {
		cookie = cookie_handler.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	if len(values) != 2 {
		http.Error(w, err.Error(), 500)
		return
	}

	uuid = values[0]

	log.Printf("IN GetURLByID controller is %T", controller)
	short := chi.URLParam(r, "articleID")
	if short == "" {
		http.Error(w, "url is empty", 404)
		return
	}

	log.Println(short)

	orig, err := controller.GetOrigByShort(r.Context(), uuid, short)
	log.Println("ORIG:", orig)
	if err != nil {
		if errors.Is(err, &repository_db.ErrItemGone{Err: err}) {
			http.Error(w, err.Error(), 410)
			return
		}
		log.Println("IN ERR HANDLER GetOrigByShort")
		http.Error(w, err.Error(), 404)
		return
	}

	if orig == "" {
		http.NotFound(w, r)
		return
	}

	log.Printf("In get_url_by_id: %s\n", orig)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Location", orig)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
