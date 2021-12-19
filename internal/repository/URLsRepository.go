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
	urls []models.URLs
	file *os.File
}



func New(filename string) *URLsRepository {
	s := &URLsRepository{}
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

func (repo *URLsRepository) StoreURL(uuid string, orig string) (string, error) {
	for {
		short := key_gen.KeyGen()
		if ok := repo.set(uuid, short, orig); ok {
			if err := repo.save(uuid, short, orig); err != nil {
				log.Println("error saving to URLStore:", err)
				return "", err
			}
			return short, nil
		}
	}
}

func (repo *URLsRepository) GetURLsList(uuid, baseURL string) []models.URLs {
	var user []models.URLs
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, v := range repo.urls {
		if v.UUID == uuid {
			user = append(user, models.URLs{
				ShortURL: baseURL + "/" + v.ShortURL,
				OriginalURL: v.OriginalURL,
			})
		}
	}
	return user
}

func (repo *URLsRepository) URLsDetail(short string) (string, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	log.Println(repo.urls)
	for _, v := range repo.urls {
		if v.ShortURL == short {
			return v.OriginalURL, nil
		}
	}
	return "", errors.New("URL not found")
}

func (repo *URLsRepository) save(uuid string, short, orig string) error {
	e := json.NewEncoder(repo.file)
	return e.Encode(models.URLs{UUID: uuid, ShortURL: short, OriginalURL: orig})
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
			repo.set(r.UUID, r.ShortURL, r.OriginalURL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (repo *URLsRepository) set(uuid string, short, orig string) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.urls = append(repo.urls, models.URLs{UUID: uuid, ShortURL: short, OriginalURL: orig})
	return true
}

