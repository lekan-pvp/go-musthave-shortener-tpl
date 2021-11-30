package shortener


type BodyRequest struct {
	URL string `json:"url"`
}

type BodyResponse struct {
	Result string `json:"result"`
}