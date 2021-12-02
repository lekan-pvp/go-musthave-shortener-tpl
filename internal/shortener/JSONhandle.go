package shortener


type BodyRequest struct {
	GoalURL string `json:"url"`
}

type BodyResponse struct {
	ResultURL string `json:"result"`
}
