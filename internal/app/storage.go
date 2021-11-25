package app

type Store struct {
	db map[string]string
}

type Storager interface {
	GetById(uuid string) string
	PutURL(url string) string
}

func NewStore() *Store  {
	store := &Store{db: map[string]string{}}
	return store
}

func (s Store) Get(uuid string) string {

	id := prefix + uuid
	return s.db[id]
}

func (s Store) Put(URL string) string{
	id := Shorting()
	s.db[id] = URL
	return id
}
