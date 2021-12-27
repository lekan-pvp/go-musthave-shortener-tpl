package interfaces

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/config"
	"github.com/go-musthave-shortener-tpl/internal/models"
)

type Storager interface {
	New(cfg *config.Config)
	InsertUserRepo(ctx context.Context, userID string, shortURL string, origURL string) (string, error)
	GetOrigByShortRepo(ctx context.Context, shortURL string) (string, error)
	GetURLsListRepo(ctx context.Context, uuid string) ([]models.URLs, error)
	CheckPingRepo(ctx context.Context) error
	BanchApiRepo(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error)
}
