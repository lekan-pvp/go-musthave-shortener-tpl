package controllers

import (
	"encoding/json"
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"io"
	"log"
	"net/http"
	"strings"
)

type short struct {
	Key string `json:"result"`
}

type long struct {
	Url string	`json:"url"`
}



func (controller *URLsController) APIShorten(w http.ResponseWriter, r *http.Request) {
	short := short{}
	long := long{}

	var uuid string
	var cookie *http.Cookie
	var err error

	cookie, err = r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie) {
		log.Println(err)
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

	if err := json.Unmarshal(body, &long); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	shortURL, err := controller.CreateURL(uuid, long.Url)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	short.Key = controller.Cfg.BaseURL + "/" + shortURL
	result, err := json.Marshal(&short)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(result))
}
