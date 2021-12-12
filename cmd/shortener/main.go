package main

import (
	"github.com/go-musthave-shortener-tpl/internal/config"
	"log"
	"net/http"
)

func main() {

	var cfg = config.New()

	log.Println("creating router...")

	router := ChiRouter().InitRouter()

	log.Println("start application")
	log.Println("server is listening port", cfg.ServerAddress)

	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
