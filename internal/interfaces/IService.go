package interfaces

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/models"
)

type IService interface {
	CreateURL(uuid string, orig string) (string, error)
	GetURLs(short string) (string, error)
	ListByUUID(uuid, baseURL string) []models.URLs
	PingDB(ctx context.Context) error
	CloseDB() error
	InsertUserDB(ctx context.Context, userID string, shortURL string, origURL string) error
	GetOrigByShortDB(ctx context.Context, shortURL string) (string, error)
	GetURLsListDB(ctx context.Context, uuid string) ([]models.URLs, error)
}
