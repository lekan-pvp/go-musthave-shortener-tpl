package interfaces

import "github.com/go-musthave-shortener-tpl/internal/models"

type IURLRepository interface {
	StoreURL(uuid string, orig string) (string, error)
	URLsDetail(url string) (string, error)
	GetURLsList(uuid, baseURL string) []models.URLs
}
