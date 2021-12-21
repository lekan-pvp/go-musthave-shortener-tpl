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
type ResultSlice []*URLS

func (controller *Controller) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	var resultSlice *ResultSlice
	cookie, err := r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie){
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(204)
		return
	}

	values := strings.Split(cookie.Value, ":")
	uuid := values[0]

	out = controller.ListByUUID(uuid, controller.Cfg.BaseURL)
	resultSlice = resultList(out)

	log.Println(resultSlice)

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

	w.WriteHeader(200)

	w.Write(marshaled)
}

func resultList(out []models.URLs) *ResultSlice {
	var result ResultSlice
	for _, v := range out {
		result = append(result, &URLS{ShortURL: v.ShortURL, OriginalURL: v.OriginalURL})
	}
	return &result
}

//func (u *URLS) MarshalJSON() ([]byte, error) {
//	arr := []interface{}{u.ShortURL, u.OriginalURL}
//	return json.Marshal(arr)
//}
//
//func (u *URLS) UnmarshalJSON(bs []byte) error {
//	arr := []interface{}{}
//	err := json.Unmarshal(bs, &arr)
//	if err != nil {
//		return err
//	}
//	u.ShortURL = arr[0].(string)
//	u.OriginalURL = arr[1].(string)
//	return nil
//}
