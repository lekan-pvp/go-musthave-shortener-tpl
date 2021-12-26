package service_memory

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
	"github.com/go-musthave-shortener-tpl/internal/models"
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
