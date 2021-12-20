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
var resultSlice []URLS

func (controller *URLsController) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil || cookie_handler.CheckCookie(cookie){
		log.Fatalf("From GetUserURLs %v", err)
		cookie = cookie_handler.CreateCookie()
		http.SetCookie(w, cookie)
	}

	values := strings.Split(cookie.Value, ":")
	uuid := values[0]

	log.Printf("%s", uuid)

	out = controller.GetURLsListByUUID(uuid, controller.Cfg.BaseURL)
	resultSlice = controller.resultList(out)

	log.Println("OUT:", out)
	log.Println("RESULT:", resultSlice)

	result, err := json.MarshalIndent(&resultSlice, "", " ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}



	w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if len(out) == 0 {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(200)
	}

	w.Write(result)
}

func (controller *URLsController) resultList(out []models.URLs) []URLS {
	var result []URLS
	for _, v := range out {
		result = append(result, URLS{ShortURL: v.ShortURL, OriginalURL: v.OriginalURL})
	}
	return result
}
