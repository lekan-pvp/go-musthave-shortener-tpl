package controllers

import (
	"github.com/go-musthave-shortener-tpl/internal/config"
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
)

type Controller struct {
	interfaces.IService
	Cfg *config.Config
}
