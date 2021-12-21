package controllers

import (
	"bytes"
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
type ResultSlice []URLS

func (controller *Controller) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	var resultSlice  = New()
	cookie, err := r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie){
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}

	values := strings.Split(cookie.Value, ":")
	uuid := values[0]

	out = controller.ListByUUID(uuid, controller.Cfg.BaseURL)
	resultSlice.Add(out)

	if resultSlice == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}


	w.Header().Set("Content-Type", "application/json; charset=utf-8")



	marshaled, err := json.Marshal(resultSlice)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("%s", marshaled)

	w.WriteHeader(200)

	var buf bytes.Buffer



	fmt.Fprint(w, marshaled)
}

func (r *ResultSlice) Add(out []models.URLs) {
	for _, v := range out {
		pnew := URLS{
			ShortURL: v.ShortURL,
			OriginalURL: v.OriginalURL,
		}
		*r = append(*r, pnew)
	}
}

func New() *ResultSlice {
	var arr ResultSlice
	return &arr
}

