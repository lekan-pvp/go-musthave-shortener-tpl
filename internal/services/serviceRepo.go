package services

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/models"
)


func (service *Service) CreateURL(uuid string, orig string) (string, error) {
	result, err := service.StoreURL(uuid, orig)
	return result, err
}

func (service *Service) GetURLs(short string) (string, error) {
	result, err := service.URLsDetail(short)
	return result, err
}

func (service *Service) ListByUUID(uuid, baseURL string) []models.URLs {
	result := service.GetURLsList(uuid, baseURL)
	return result
}

func (service *Service) PingDB(ctx context.Context) error {
	if err := service.CheckPingDB(ctx); err != nil {
		return err
	}
	return nil
}

func (service *Service) CloseDB() error {
	err := service.CloseDBRepo()
	if err != nil {
		return err
	}
	return nil
}
