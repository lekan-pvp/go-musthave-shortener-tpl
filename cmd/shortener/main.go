package main

import (
	"fmt"
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

	handler := shortener.NewHandler(cfg)
	handler.Register(router)

	log.Println("start application")
	log.Println("server is listening port", cfg.ServerAddress)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.ServerAddress, "8080"), router))
}



