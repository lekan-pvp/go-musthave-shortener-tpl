package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL string `env:"BASE_URL"`
}

var cfg *Config
//var once sync.Once

func GetConfig() *Config {
	//once.Do(func() {
	//	log.Println("read application configuration")
	//
	//	cfg = &Config{}
	//
	//	err := env.Parse(cfg)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println(cfg)
	//})
	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", cfg)
	return cfg
}
