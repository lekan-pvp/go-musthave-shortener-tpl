package interfaces

type IURLRepository interface {
	StoreURL(key string) (string, error)
	URLsDetail(url string) (string, error)
}
