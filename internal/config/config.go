package config

import (
	"github.com/caarlos0/env/v6"
	"log"
	"sync"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL string `env:"BASE_URL"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := env.Parse(instance); err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v", instance)
	})
	return instance
}
