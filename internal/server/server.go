package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-musthave-shortener-tpl/internal/config"
	"github.com/go-musthave-shortener-tpl/internal/controllers"
	"github.com/go-musthave-shortener-tpl/internal/mware"
	"github.com/go-musthave-shortener-tpl/internal/repository"
	"github.com/go-musthave-shortener-tpl/internal/services"
	"log"
	"net/http"
)

func Run() {
	cfg := config.New()

	repo := repository.New(cfg.FileStoragePath, cfg.DatabaseDSN)
	service := &services.Service{repo}
	controller := controllers.Controller{service, cfg}

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/user/urls", controller.GetUserURLs)
	router.Get("/ping", controller.PingDBHandler)
	router.With(mware.RequestHandle, mware.GzipHandle).Post("/", controller.AddURL)
	router.Route("/{articleID:[a-zA-Z]+}", func(r chi.Router) {
		r.Use(mware.GzipHandle)
		r.Get("/", controller.GetURLByID)
	})
	router.With(mware.RequestHandle, mware.GzipHandle).Post("/api/shorten", controller.APIShorten)



	log.Println("creating router...")
	log.Println("start application")
	log.Println("server is listening port", cfg.ServerAddress)

	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
