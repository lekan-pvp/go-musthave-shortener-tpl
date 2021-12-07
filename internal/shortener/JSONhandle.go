package shortener

type BodyRequest struct {
	LongURL string `json:"url"`
}

type BodyResponse struct {
	ShortURL string `json:"result"`
}

