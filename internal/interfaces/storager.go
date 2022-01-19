package interfaces

import (
	"context"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
	"sync"
)

type Storager interface {
	New(cfg *config.Config)
	InsertUserRepo(ctx context.Context, userID string, shortURL string, origURL string) (string, error)
	GetOrigByShortRepo(ctx context.Context, shortURL string) (string, error)
	GetURLsListRepo(ctx context.Context, uuid string) ([]models.URLs, error)
	CheckPingRepo(ctx context.Context) error
	BanchApiRepo(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error)
	UpdateURLsRepo(ctx context.Context, uuid string, shortURLs []string) error
	DeleteURLsRepo(ctx context.Context, uuid string, short string, errCh chan<- error, wg *sync.WaitGroup)
}
