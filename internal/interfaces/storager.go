package interfaces

import (
	"context"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
)

type Storager interface {
	New(cfg *config.Config)
	InsertUserRepo(ctx context.Context, userID string, shortURL string, origURL string) (string, error)
	GetOrigByShortRepo(ctx context.Context, shortURL string) (*models.OriginLink, error)
	GetURLsListRepo(ctx context.Context, uuid string) ([]models.URLs, error)
	CheckPingRepo(ctx context.Context) error
	BanchAPIRepo(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error)
	UpdateURLsRepo(ctx context.Context, shortURLs []string) error
}
