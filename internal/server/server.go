package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-musthave-shortener-tpl/internal/config"
	"github.com/go-musthave-shortener-tpl/internal/controllers"
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
	"github.com/go-musthave-shortener-tpl/internal/mware"
	repository_db "github.com/go-musthave-shortener-tpl/internal/repository-db"
	repository_memory "github.com/go-musthave-shortener-tpl/internal/repository-memory"
	"github.com/go-musthave-shortener-tpl/internal/service-memory"
	"log"
	"net/http"
)


func Run() {
	cfg := config.New()

	repo := New(cfg)

	repo.New(cfg)

	service := &service_memory.Service{repo}
	controller := controllers.Controller{service, cfg}

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/user/urls", controller.GetUserURLs)
	router.Get("/ping", controller.PingDBHandler)
	router.With(mware.RequestHandle, mware.GzipHandle).Post("/", controller.AddURL)
	router.With(mware.GzipHandle).Get("/{articleID}", controller.GetURLByID)
	router.With(mware.RequestHandle, mware.GzipHandle).Post("/api/shorten", controller.APIShorten)



	log.Println("creating router...")
	log.Println("start application")
	log.Println("server is listening port", cfg.ServerAddress)
	log.Println("db connect at", cfg.DatabaseDSN)

	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}

func New(cfg *config.Config) interfaces.Storager {
	if cfg.FileStoragePath != "" {
		return &repository_memory.MemoryRepository{}
	}
	if cfg.DatabaseDSN != "" {
		return &repository_db.DBRepository{}
	}
	return nil
}
