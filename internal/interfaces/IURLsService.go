package interfaces

type IURLsService interface {
	CreateURL(url string) (string, error)
	GetURLsDetail(key string) (string, error)
}
