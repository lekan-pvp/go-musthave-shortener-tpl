package models

type URLs struct {
	UUID string `json:"-"`
	ShortURL string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

