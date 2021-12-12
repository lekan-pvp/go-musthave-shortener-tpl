package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"short_urls.json"`
}

var cfg Config

func init() {
	log.Println("init configuration...")
	cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&cfg.ServerAddress, "a", "localhost:8080", "адрес запуска HTTP-сервера")
	flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "базовый адрес результирующего сокращённого URL")
	flag.StringVar(&cfg.FileStoragePath, "f", "short_url.json", "путь к файлу с сокращёнными URL")

	flag.Parse()
}

func New() *Config {
	return &cfg
}
