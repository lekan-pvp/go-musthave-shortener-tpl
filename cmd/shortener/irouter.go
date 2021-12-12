package main

import "github.com/go-chi/chi"

type IChiRouter interface {
	InitRouter() *chi.Mux
}
