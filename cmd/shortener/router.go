package main

import (
	"github.com/go-chi/chi"
	"sync"
)

type router struct {}

func (r *router) InitRouter() *chi.Mux {
	urlController := ServiceContainer().InjectURLController()

	router := chi.NewRouter()

	router.Post("/", urlController.AddURL)
	router.Get("/{articleID}", urlController.GetURLByID)
	router.Post("/api/shorten", urlController.APIShorten)

	return router
}

var (
	m *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}

