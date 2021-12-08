package shortener

import (
	"sync"
)

type Store struct {
	mx sync.Mutex
	db map[string]string
}

type Storer interface {
	Get(string) string
	Put(string) string
}

func NewStore() *Store  {
	return &Store{
		db: make(map[string]string),
	}
}

func (s *Store) Get(uuid string) (string, bool) {
	s.mx.Lock()
	defer s.mx.Unlock()
	val, ok := s.db[uuid]
	return val, ok
}

func (s *Store) Put(URL string) string{
	s.mx.Lock()
	defer s.mx.Unlock()
	id := Shorting()
	s.db[id] = URL
	return id
}
