package servicedb

import (
	"context"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/interfaces"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
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

func (service *Service) GetOrigByShort(ctx context.Context, shortURL string) (*models.OriginLink, error) {
	result := &models.OriginLink{}
	result, err := service.GetOrigByShortRepo(ctx, shortURL)
	if err != nil {
		return nil, err
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

func (service *Service) BanchAPI(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error) {
	result, err := service.BanchAPIRepo(ctx, uuid, in, shortBase)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *Service) UpdateURLs(ctx context.Context, shortBases []string) error {
	err := service.UpdateURLsRepo(ctx, shortBases)
	if err != nil {
		return err
	}
	return nil
}
