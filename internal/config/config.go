package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"short_url.json"`
	DatabaseDSN 	string `env:"DATABASE_DSN" envDefault:"user=postgres password='postgres' dbname=pqgotest sslmode=disable"`
}

var instance *Config

func New() *Config {
	log.Println("init config...")
	instance = &Config{}
	if err := env.Parse(instance); err != nil {
		log.Fatal(err)
	}

	serverAddressPtr := flag.String("a", instance.ServerAddress, "адрес запуска HTTP-сервера")
	baseURLPtr := flag.String("b", instance.BaseURL, "базовый адрес результирующего сокращённого URL")
	fileStoragePathPtr := flag.String("f", instance.FileStoragePath, "путь до файла с сокращёнными URL")
	databaseDSN := flag.String("d", instance.DatabaseDSN, "адрес подключения к БД")

	flag.Parse()



	instance.ServerAddress = *serverAddressPtr
	instance.BaseURL = *baseURLPtr

	if fileStoragePathPtr == nil {
		instance.DatabaseDSN = *databaseDSN
	} else {
		instance.FileStoragePath = *fileStoragePathPtr
	}

	return instance
}
