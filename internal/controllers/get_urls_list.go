package controllers

import (
	"encoding/json"
	"fmt"
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
var resultSlice []URLS

func (controller *Controller) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	log.Println("IN GetUserURLs:")

	cookie, err := r.Cookie("token")
	if err != nil || cookie_handler.CheckCookie(cookie){
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}

	values := strings.Split(cookie.Value, ":")
	uuid := values[0]

	out = controller.ListByUUID(uuid, controller.Cfg.BaseURL)
	resultSlice = controller.resultList(out)

	if len(resultSlice) == 0 {
		log.Println("len(resultSlice)=", len(resultSlice))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}

	result, err := json.MarshalIndent(resultSlice, "", " ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", result)
	w.Write(result)
}

func (controller *Controller) resultList(out []models.URLs) []URLS {
	var result []URLS
	for _, v := range out {
		result = append(result, URLS{ShortURL: v.ShortURL, OriginalURL: v.OriginalURL})
	}
	return result
}
