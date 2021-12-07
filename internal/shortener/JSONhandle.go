package shortener

type BodyRequest struct {
	LongURL string `json:"long-url"`
}

type BodyResponse struct {
	ShortURL string `json:"short-url"`
}

