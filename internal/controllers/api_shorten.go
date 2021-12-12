package controllers

import (
	"encoding/json"
	"io"
	"net/http"
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
	key, err := controller.CreateURL(long.Url)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	short.Key = controller.Cfg.BaseURL + "/" + key
	result, err := json.Marshal(&short)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(result))
}
