package storage

import (
	"encoding/json"
	"errors"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/shortener/gen_key"
	"io"
	"log"
	"os"
	"sync"
)

type URLStore struct {
	mu sync.RWMutex
	urls map[string]string
	file *os.File
}

type record struct {
	Key, URL string
}

// NewStore -- фабрика хранилища
func NewStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal("URLStore:", err)
	}
	s.file = f
	if err := s.load(); err != nil {
		log.Println("Error loading data in URLStore:", err)
	}
	return s
}

// save -- запись в файл
func (s *URLStore) save(key, url string) error {
	e := json.NewEncoder(s.file)
	return e.Encode(record{key, url})
}

// load -- чтение из файла
func (s *URLStore) load() error {
	if _, err := s.file.Seek(0, 0); err != nil {
		return err
	}
	d := json.NewDecoder(s.file)
	var err error
	for err == nil {
		var r record
		if err = d.Decode(&r); err != nil {
			s.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

// Get -- достает из хранилища
func (s *URLStore) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.Unlock()

	url, ok := s.urls[key]
	if !ok {
		return "", errors.New("short URL not found ")
	}
	return url, nil
}

// Set -- записывает в хранилище
func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

// Put -- генерирует короткий URL, записывает в хранилище и в файл
func (s *URLStore) Put(url string) string{
	for {
		key := gen_key.GenKey()
		if ok := s.Set(key, url); ok {
			if err := s.save(key, url); err != nil {
				log.Println("error saving to URLStore:", err)
			}
			return key
		}
	}
}
