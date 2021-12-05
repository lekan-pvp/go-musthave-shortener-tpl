package shortener

import (
	"errors"
	"log"
	"sync"
)

type MemoryStore struct {
	mx sync.Mutex
	db map[string]string
}

func NewStore() *MemoryStore  {
	return &MemoryStore{
		db: make(map[string]string),
	}
}

func (s *MemoryStore) Get(uuid string) (string, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	val, ok := s.db[uuid]
	if !ok {
		return "", errors.New("short URL not found ")
	}
	return val, nil
}

func (s *MemoryStore) Put(URL string, shortURL string) string{
	log.Println(shortURL)
	s.mx.Lock()
	defer s.mx.Unlock()
	s.db[shortURL] = URL
	return shortURL
}
