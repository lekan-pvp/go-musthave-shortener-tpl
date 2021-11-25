package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/app"
	"log"
	"net/http"

)

func main() {
	router := chi.NewRouter()

	handler := app.NewHandler()
	handler.Register(router)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}



