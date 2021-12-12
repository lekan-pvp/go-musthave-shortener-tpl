package repository

import (
	"encoding/json"
	"errors"
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
	"github.com/go-musthave-shortener-tpl/internal/key_gen"
	"github.com/go-musthave-shortener-tpl/internal/models"
	"io"
	"log"
	"os"
	"sync"
)

type URLsRepository struct {
	URLRepository interfaces.IURLRepository
	mu sync.RWMutex
	urls map[string]string
	file *os.File
}



func New(filename string) *URLsRepository {
	s := &URLsRepository{urls: make(map[string]string)}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("error loading in Store:", err)
	}
	s.file = f
	if err := s.load(); err != nil {
		log.Println("error loading data in Store:", err)
	}
	return s
}

func (repo *URLsRepository) StoreURL(url string) (string, error) {
	for {
		key := key_gen.KeyGen()
		if ok := repo.set(key, url); ok {
			if err := repo.save(key, url); err != nil {
				log.Println("error saving to URLStore:", err)
				return "", err
			}
			return key, nil
		}
	}
}


func (repo *URLsRepository) URLsDetail(key string) (string, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	log.Println(repo.urls)
	url, ok := repo.urls[key]
	if !ok {
		return "", errors.New("short URL not found")
	}
	return url, nil
}

func (repo *URLsRepository) save(key, url string) error {
	e := json.NewEncoder(repo.file)
	return e.Encode(models.URLs{Key: key, URL: url})
}

func (repo *URLsRepository) load() error  {
	if _, err := repo.file.Seek(0, 0); err != nil {
		return err
	}
	d := json.NewDecoder(repo.file)
	var err error
	for err == nil {
		var r models.URLs
		if err = d.Decode(&r); err == nil {
			repo.set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (repo *URLsRepository) set(key, url string) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if _, present := repo.urls[key]; present {
		return false
	}
	repo.urls[key] = url
	return true
}



