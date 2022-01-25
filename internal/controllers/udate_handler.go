package controllers

import (
	"encoding/json"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/cookieserver"
	"io"
	"log"
	"net/http"
	"strings"
)

func (service *Controller) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var in []string
	cookie, err := r.Cookie("token")
	if err != nil || !cookieserver.CheckCookie(cookie) {
		cookie = cookieserver.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	if len(values) != 2 {
		log.Panicln("cookie format error...")
		http.Error(w, err.Error(), 500)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("reading body error...")
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(body, &in)
	if err != nil {
		log.Println("decoding json error...")
		http.Error(w, err.Error(), 500)
		return
	}

	err = service.UpdateURLs(r.Context(), in)
	if err != nil {
		log.Println("update db error")
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(202)
}
