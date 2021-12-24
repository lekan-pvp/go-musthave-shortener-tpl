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
	CreateTableDBRepo(ctx context.Context, name string) error
	InsertUserDBRepo(ctx context.Context,tabname string, userID string, shortURL string, origURL string) error
}
