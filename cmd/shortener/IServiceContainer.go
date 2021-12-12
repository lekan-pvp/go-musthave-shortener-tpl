package main

import "github.com/go-musthave-shortener-tpl/internal/controllers"

type IServiceContainer interface {
	InjectURLController() controllers.URLsController
}




