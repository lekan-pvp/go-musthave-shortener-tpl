package services

import (
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
)

type URLsService struct {
	interfaces.IURLRepository
}

func (service *URLsService) CreateURL(url string) (string, error) {
	result, err := service.StoreURL(url)
	return result, err
}

func (service *URLsService) GetURLs(key string) (string, error) {
	result, err := service.URLsDetail(key)
	return result, err
}
