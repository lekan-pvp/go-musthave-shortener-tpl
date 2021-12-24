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

	cookie, err = r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie){
		cookie = cookie_handler.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	uuid = values[0]

	err = controller.CreateTableDB(r.Context(), "users")
	if err != nil {
		log.Println("error table create")
		http.Error(w, err.Error(), 500)
		return
	} else {
		log.Println("Table created")
	}

	defer func() {
		err := controller.CloseDB()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

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

	err = controller.InsertUserDB(r.Context(), "users", uuid, short, orig)
	if err != nil {
		log.Println("error insert in DB:", err)
		http.Error(w, err.Error(), 500)
		return
	} else {
		log.Println("user inserted")
	}

	w.Write([]byte(controller.Cfg.BaseURL + "/" + short))
}
