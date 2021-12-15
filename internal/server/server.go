package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-musthave-shortener-tpl/internal/config"
	"github.com/go-musthave-shortener-tpl/internal/controllers"
	"github.com/go-musthave-shortener-tpl/internal/repository"
	"github.com/go-musthave-shortener-tpl/internal/services"
	"log"
	"net/http"
)

func Run() {
	cfg := config.New()

	urlRepo := repository.New(cfg.FileStoragePath)
	urlService := &services.URLsService{urlRepo}
	urlController := controllers.URLsController{urlService, cfg}

	router := chi.NewRouter()

	router.Use(middleware.Compress(5))

	router.Post("/", urlController.AddURL)
	router.Get("/{articleID}", urlController.GetURLByID)
	router.Post("/api/shorten", urlController.APIShorten)

	log.Println("creating router...")
	log.Println("start application")
	log.Println("server is listening port", cfg.ServerAddress)

	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
