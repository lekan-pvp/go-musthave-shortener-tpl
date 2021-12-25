package interfaces

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/models"
)

type IRepository interface {
	StoreURL(uuid string, orig string) (string, error)
	URLsDetail(url string) (string, error)
	GetURLsList(uuid, baseURL string) []models.URLs
	CheckPingDB(ctx context.Context) error
	CloseDBRepo() error
	InsertUserDBRepo(ctx context.Context, userID string, shortURL string, origURL string) error
	GetOrigByShortDBRepo(ctx context.Context, shortURL string) (string, error)
	GetURLsListDBRepo(ctx context.Context, uuid string) ([]models.URLs, error)
}
