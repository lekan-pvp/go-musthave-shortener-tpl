package services

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/models"
)

func (service *Service) InsertUserDB(ctx context.Context, userID string, shortURL string, origURL string) error {
	err := service.InsertUserDBRepo(ctx, userID, shortURL, origURL)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) GetOrigByShortDB(ctx context.Context, shortURL string) (string, error) {
	result, err := service.GetOrigByShortDBRepo(ctx, shortURL)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (service *Service) GetURLsListDB(ctx context.Context, uuid string) ([]models.URLs, error) {
	result, err := service.GetURLsListDBRepo(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}
