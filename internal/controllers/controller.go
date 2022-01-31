package controllers

import (
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/interfaces"
)

type Controller struct {
	interfaces.Servicer
	Cfg *config.Config
}
