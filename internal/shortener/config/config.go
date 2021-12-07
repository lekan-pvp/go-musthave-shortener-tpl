package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"short_urls.json"`
}

var instance *Config

func GetConfig() *Config {
	instance = &Config{}
	if err := env.Parse(instance); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", instance)

	return instance
}
