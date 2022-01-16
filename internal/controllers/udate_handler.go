package controllers

import (
	"encoding/json"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/cookie_handler"
	"io"
	"net/http"
	"strings"
)

func (controller *Controller) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var uuid string
	in := make([]string, 0)

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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(body, &in)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = controller.UpdateURLs(r.Context(), uuid, in)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(202)
}
