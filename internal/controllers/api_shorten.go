package controllers

import (
	"encoding/json"
	"github.com/jackc/pgerrcode"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/cookieServer"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/keyGen"
	"github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"strings"
)

type short struct {
	Key string `json:"result"`
}

type long struct {
	URL string `json:"url"`
}

func (controller *Controller) APIShorten(w http.ResponseWriter, r *http.Request) {
	short := short{}
	long := long{}

	var uuid string
	var cookie *http.Cookie
	var err error

	cookie, err = r.Cookie("token")
	if err != nil || !cookieServer.CheckCookie(cookie) {
		cookie = cookieServer.CreateCookie()
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

	status := http.StatusCreated

	generatedShortURL := keyGen.GenerateShortLink(long.URL, uuid)
	shortURL, err := controller.InsertUser(r.Context(), uuid, generatedShortURL, long.URL)
	if err != nil {
		if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
			status = http.StatusConflict
		} else {
			log.Println("error insert in DB:", err)
			http.Error(w, err.Error(), 500)
			return
		}
	}

	short.Key = controller.Cfg.BaseURL + "/" + shortURL
	result, err := json.Marshal(&short)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(status)
	w.Write([]byte(result))
}
