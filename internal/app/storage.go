package app

type Store struct {
	db map[string]string
}

type Storer interface {
	Get(string) string
	Put(string) string
}

func NewStore() *Store  {
	store := &Store{db: map[string]string{}}
	return store
}

func (s *Store) Get(uuid string) string {
	return s.db[uuid]
}

func (s *Store) Put(URL string) string{
	id := Shorting()
	s.db[id] = URL
	return id
}
