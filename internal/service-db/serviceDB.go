package service_db

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
	"github.com/go-musthave-shortener-tpl/internal/models"
)

type Service struct {
	interfaces.Storager
}

func (service *Service) InsertUser(ctx context.Context, userID string, shortURL string, origURL string) (string, error) {
	short, err := service.InsertUserRepo(ctx, userID, shortURL, origURL)
	if err != nil {
		return "", err
	}
	return short, nil
}

func (service *Service) GetOrigByShort(ctx context.Context, shortURL string) (string, error) {
	result, err := service.GetOrigByShortRepo(ctx, shortURL)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (service *Service) GetList(ctx context.Context, uuid string) ([]models.URLs, error) {
	result, err := service.GetURLsListRepo(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *Service) CheckPing(ctx context.Context) error {
	err := service.CheckPingRepo(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) BanchApi(ctx context.Context, in []models.BatchIn, shortBase string) []models.BatchResult {
	result := service.BanchApiRepo(ctx, in, shortBase)
	return result
}
