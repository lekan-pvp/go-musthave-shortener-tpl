package service_memory

import (
	"context"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/interfaces"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
)

type Service struct {
	interfaces.Storager
}

func (service *Service) InsertUser(ctx context.Context, userID string, shortURL string, origURL string) (string, error) {
	result, err := service.InsertUserRepo(ctx, userID, shortURL, origURL)
	return result, err
}

func (service *Service) GetOrigByShort(ctx context.Context, shortURL string) (string, error) {
	result, err := service.GetOrigByShortRepo(ctx, shortURL)
	return result, err
}

func (service *Service) GetList(ctx context.Context, uuid string) ([]models.URLs, error) {
	result, err := service.GetURLsListRepo(ctx, uuid)
	return result, err
}

func (service *Service) CheckPing(ctx context.Context) error {
	err := service.CheckPingRepo(ctx)
	return err
}

func (service *Service) BanchApi(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error) {
	result, err := service.BanchApiRepo(ctx, uuid, in, shortBase)
	return result, err
}

func (service *Service) UpdateURLs(ctx context.Context, uuid string, shortBase []string) error {
	err := service.UpdateURLsRepo(ctx, uuid, shortBase)
	return err
}
