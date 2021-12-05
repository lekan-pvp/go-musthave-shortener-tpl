package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/shortener"
	"log"
	"net/http"
)

func main() {
	log.Println("creating router...")

	router := chi.NewRouter()

	log.Println("register shorner handler...")

	cfg := config.GetConfig()

	handler := shortener.NewHandler()
	handler.Register(router)

	log.Println("start application")
	log.Println("server is listening port", cfg.ServerAddress)

	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}



