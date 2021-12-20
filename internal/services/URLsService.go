package services

import (
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
	"github.com/go-musthave-shortener-tpl/internal/models"
)

type URLsService struct {
	interfaces.IURLRepository
}

func (service *URLsService) CreateURL(uuid string, orig string) (string, error) {
	result, err := service.StoreURL(uuid, orig)
	return result, err
}

func (service *URLsService) GetURLs(short string) (string, error) {
	result, err := service.URLsDetail(short)
	return result, err
}

func (service *URLsService) ListByUUID(uuid, baseURL string) []models.URLs {
	result := service.GetURLsList(uuid, baseURL)
	return result
}
