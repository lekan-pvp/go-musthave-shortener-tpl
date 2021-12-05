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

	cfg := config.GetConfig()


	log.Println("register shorner handler...")

	handler := shortener.NewHandler(cfg)
	handler.Register(router)

	log.Println("start application")
	log.Printf("server is listening port %v", cfg.ServerAddress)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.ServerAddress,"8080"), router))
}



