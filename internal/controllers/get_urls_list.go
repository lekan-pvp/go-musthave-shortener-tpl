package controllers

import (
	"encoding/json"
	"fmt"
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
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}

	values := strings.Split(cookie.Value, ":")
	uuid := values[0]

	log.Printf("%s", uuid)

	out = controller.GetURLsListByUUID(uuid, controller.Cfg.BaseURL)
	resultSlice = controller.resultList(out)

	if len(resultSlice) == 0 {
		log.Println("len(resultSlice)=", len(resultSlice))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}


	result, err := json.Marshal(&resultSlice)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(string(result))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Write(result)
	w.WriteHeader(200)


}

func (controller *URLsController) resultList(out []models.URLs) []URLS {
	var result []URLS
	for _, v := range out {
		result = append(result, URLS{ShortURL: v.ShortURL, OriginalURL: v.OriginalURL})
	}
	return result
}
