package interfaces

import (
	"context"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
)

type Servicer interface {
	InsertUser(ctx context.Context, userID string, shortURL string, origURL string) (string, error)
	GetOrigByShort(ctx context.Context, shortURL string) (string, error)
	GetList(ctx context.Context, uuid string) ([]models.URLs, error)
	CheckPing(ctx context.Context) error
	BanchAPI(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error)
	UpdateURLs(ctx context.Context, shortBase []string) error
	//DeleteURLs(ctx context.Context, uuid string, short string, errCh chan<- error, wg *sync.WaitGroup)
	//DeleteItem(ctx context.Context, short string) error
}
