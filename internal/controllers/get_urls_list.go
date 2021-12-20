package controllers

import (
	"encoding/json"
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"github.com/go-musthave-shortener-tpl/internal/models"
	"log"
	"net/http"
	"strings"
)



type URLS struct {
	ShortURL string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

var out []models.URLs
var ResultSlice []URLS

func (controller *Controller) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie){
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}

	values := strings.Split(cookie.Value, ":")
	uuid := values[0]

	out = controller.ListByUUID(uuid, controller.Cfg.BaseURL)
	log.Println("OUT=", out)
	ResultSlice = controller.resultList(out)
	log.Println("resultSlice=", ResultSlice)

	if len(ResultSlice) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}

	result, err := json.MarshalIndent(&ResultSlice, "", " ")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	log.Printf("%s\n", result)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	w.Write(result)
}

func (controller *Controller) resultList(out []models.URLs) []URLS {
	var result []URLS
	for _, v := range out {
		result = append(result, URLS{ShortURL: v.ShortURL, OriginalURL: v.OriginalURL})
	}
	return result
}
