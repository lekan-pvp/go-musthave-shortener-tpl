package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/controllers"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/interfaces"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/mware"
	repository_db "github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/repositoryDB"
	repository_memory "github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/repositoryMemory"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/serviceMemory"
	"log"
	"net/http"
)

func Run() {
	cfg := config.New()

	repo := New(cfg)

	repo.New(cfg)

	service := &serviceMemory.Service{repo}
	controller := controllers.Controller{service, cfg}

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/user/urls", controller.GetUserURLs)
	router.Get("/ping", controller.PingDBHandler)
	router.With(mware.RequestHandle, mware.GzipHandle).Post("/", controller.AddURL)
	router.With(mware.GzipHandle).Get("/{articleID}", controller.GetURLByID)
	router.Route("/api/shorten", func(r chi.Router) {
		r.With(mware.RequestHandle, mware.GzipHandle).Post("/", controller.APIShorten)
		r.Post("/batch", controller.APIShortenBatch)
	})
	router.Route("/api/user", func(r chi.Router) {
		r.Delete("/urls", controller.UpdateHandler)
	})

	log.Println("creating router...")
	log.Println("start application")
	log.Println("server is listening port", cfg.ServerAddress)
	log.Println("db connect at", cfg.DatabaseDSN)

	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}

func New(cfg *config.Config) interfaces.Storager {
	log.Println("cfg.FileStoragePath =", cfg.FileStoragePath)
	if cfg.FileStoragePath != "" {
		return &repository_memory.MemoryRepository{}
	}
	log.Println("cfg.DatabaseDSN =", cfg.DatabaseDSN)
	if cfg.DatabaseDSN != "" {
		return &repository_db.DBRepository{}
	}
	return nil
}
