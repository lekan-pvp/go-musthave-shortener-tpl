package interfaces

import (
	"context"
	"github.com/go-musthave-shortener-tpl/internal/models"
)

type Servicer interface {
	InsertUser(ctx context.Context, userID string, shortURL string, origURL string) (string, error)
	GetOrigByShort(ctx context.Context, shortURL string) (string, error)
	GetList(ctx context.Context, uuid string) ([]models.URLs, error)
	CheckPing(ctx context.Context) error
	BanchApi(ctx context.Context, in []models.BatchIn, shortBase string) []models.BatchResult
}
