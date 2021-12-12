package main

import (
	"github.com/go-musthave-shortener-tpl/internal/config"
	"github.com/go-musthave-shortener-tpl/internal/controllers"
	"github.com/go-musthave-shortener-tpl/internal/repository"
	"github.com/go-musthave-shortener-tpl/internal/services"
	"sync"
)

type kernel struct {

}

var (
	k *kernel
	containerOnce sync.Once
)

func (k *kernel) InjectURLController() controllers.URLsController {
	cfg := config.New()
	URLRepository := repository.New(cfg.FileStoragePath)
	URLService := &services.URLsService{URLRepository}
	URLController := controllers.URLsController{URLService, cfg}
	return URLController
}

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}