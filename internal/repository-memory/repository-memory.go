package repository_memory

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/interfaces"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/key_gen"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
	"sync"
)

type MemoryRepository struct {
	interfaces.Storager
	mu    sync.RWMutex
	users []models.URLs
	File  *os.File
}

func (s *MemoryRepository) New(cfg *config.Config) {
	f, err := os.OpenFile(cfg.FileStoragePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("error loading in Store:", err)
	}
	s.File = f
	if err := s.load(); err != nil {
		log.Println("error loading data in Store:", err)
	}
}

func (s *MemoryRepository) InsertUserRepo(ctx context.Context, userID string, shortURL string, origURL string) (string, error) {
	log.Println("IN MEM: InsertUserRepo")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users = append(s.users, models.URLs{UUID: userID, ShortURL: shortURL, OriginalURL: origURL})
	if err := s.save(userID, shortURL, origURL); err != nil {
		log.Println("error saving to URLStore:", err)
		return "", err
	}
	return shortURL, nil
}

func (s *MemoryRepository) GetURLsListRepo(ctx context.Context, uuid string) ([]models.URLs, error) {
	log.Println("IN MEM: GetURLsListRepo")
	s.mu.RLock()
	defer s.mu.RUnlock()

	var user []models.URLs

	log.Println("From GetList: ")
	log.Println(s.users)

	for _, v := range s.users {
		if v.UUID == uuid {
			user = append(user, models.URLs{
				UUID:        v.UUID,
				ShortURL:    v.ShortURL,
				OriginalURL: v.OriginalURL,
			})
		}
	}
	return user, nil
}

func (s *MemoryRepository) GetOrigByShortRepo(ctx context.Context, short string) (string, error) {
	log.Println("IN MEM: GetOrigByShortRepo")
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.users {
		if v.ShortURL == short {
			return v.OriginalURL, nil
		}
	}
	return "", errors.New("short URL not found")
}

func (s *MemoryRepository) save(uuid string, short, orig string) error {
	e := json.NewEncoder(s.File)
	return e.Encode(models.URLs{UUID: uuid, ShortURL: short, OriginalURL: orig})
}

func (s *MemoryRepository) load() error {
	if _, err := s.File.Seek(0, 0); err != nil {
		return err
	}
	d := json.NewDecoder(s.File)
	var err error
	for err == nil {
		var r models.URLs
		if err = d.Decode(&r); err == nil {
			s.setUser(r.UUID, r.ShortURL, r.OriginalURL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *MemoryRepository) setUser(uuid, short, orig string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users = append(s.users, models.URLs{UUID: uuid, ShortURL: short, OriginalURL: orig})
}

func (s *MemoryRepository) CheckPingRepo(ctx context.Context) error {
	return nil
}

func (s *MemoryRepository) BanchApiRepo(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error) {
	log.Println("BanchApiRepo IN MEMORY:")
	result := make([]models.BatchResult, 0)
	for _, v := range in {
		short := key_gen.GenerateShortLink(v.OriginalURL, v.CorrelationID)
		result = append(result, models.BatchResult{CorrelationID: v.CorrelationID, ShortURL: shortBase + "/" + short})
		s.users = append(s.users, models.URLs{UUID: uuid, ShortURL: short, OriginalURL: v.OriginalURL, CorrelationID: v.CorrelationID})
	}
	return result, nil
}

func (s *MemoryRepository) UpdateURLsRepo(ctx context.Context, uuid string, shortURLs []string) error {
	return nil
}
