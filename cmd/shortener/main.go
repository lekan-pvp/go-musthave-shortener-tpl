package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/shortener"
	"log"
	"net/http"
)

func main() {
	log.Println("creating router...")

	router := chi.NewRouter()

	log.Println("register shorner handler...")

	handler := shortener.NewHandler()
	handler.Register(router)

	log.Println("start application")
	log.Println("server is listening port 127.0.0.1:8080")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}



