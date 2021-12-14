package interfaces

type IURLsService interface {
	CreateURL(url string) (string, error)
	GetURLs(key string) (string, error)
}
