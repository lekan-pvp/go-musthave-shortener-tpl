package interfaces

import "github.com/go-musthave-shortener-tpl/internal/models"

type IURLsService interface {
	CreateURL(uuid string, orig string) (string, error)
	GetURLs(short string) (string, error)
	GetURLsListByUUID(uuid string) []models.URLs
}
