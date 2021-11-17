package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func BodyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(b))
}

func GetHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func main() {
	router := httprouter.New()
	router.POST("/", BodyHandler)
	router.GET("/:id", GetHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
