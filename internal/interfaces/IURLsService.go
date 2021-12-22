package interfaces

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/models"
)

type IURLsService interface {
	CreateURL(uuid string, orig string) (string, error)
	GetURLs(short string) (string, error)
	ListByUUID(uuid, baseURL string) []models.URLs
	PingDB(ctx context.Context) error
	CloseDB() error
}
