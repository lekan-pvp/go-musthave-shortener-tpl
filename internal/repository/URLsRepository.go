package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-musthave-shortener-tpl/internal/interfaces"
	"github.com/go-musthave-shortener-tpl/internal/key_gen"
	"github.com/go-musthave-shortener-tpl/internal/models"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
	"sync"
)

type Repository struct {
	interfaces.IRepository
	mu sync.RWMutex
	users []models.URLs
	urls map[string]string
	file *os.File
	DB *sql.DB
}

func New(filename, connStr string) *Repository {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to DB:", err)
	}

	s := &Repository{
		urls: make(map[string]string),
		DB: db,
	}
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

func (repo *Repository) StoreURL(uuid string, orig string) (string, error) {
	for {
		short := key_gen.KeyGen()
		if ok := repo.set(short, orig); ok {
			repo.setUser(uuid, short, orig)
			if err := repo.save(uuid, short, orig); err != nil {
				log.Println("error saving to URLStore:", err)
				return "", err
			}
			return short, nil
		}
	}
}

func (repo *Repository) GetURLsList(uuid, baseURL string) []models.URLs {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var user []models.URLs

	log.Println("From GetURLsList: ")
	log.Println(repo.users)

	for _, v := range repo.users {
		if v.UUID == uuid {
			user = append(user, models.URLs{
				UUID: v.UUID,
				ShortURL: baseURL + "/" + v.ShortURL,
				OriginalURL: v.OriginalURL,
			})
		}
	}
	return user
}

func (repo *Repository) URLsDetail(short string) (string, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	log.Println("From URLsDetail:")
	url, ok := repo.urls[short]
	if !ok {
		return "", errors.New("short URL not found")
	}
	return url, nil
}

func (repo *Repository) save(uuid string, short, orig string) error {
	e := json.NewEncoder(repo.file)
	return e.Encode(models.URLs{UUID: uuid, ShortURL: short, OriginalURL: orig})
}

func (repo *Repository) load() error  {
	if _, err := repo.file.Seek(0, 0); err != nil {
		return err
	}
	d := json.NewDecoder(repo.file)
	var err error
	for err == nil {
		var r models.URLs
		if err = d.Decode(&r); err == nil {
			repo.setUser(r.UUID, r.ShortURL, r.OriginalURL)
			repo.set(r.ShortURL, r.OriginalURL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (repo *Repository) setUser(uuid, short, orig string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.users = append(repo.users, models.URLs{UUID: uuid, ShortURL: short, OriginalURL: orig})
}

func (repo *Repository) set(short, orig string) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if _, present := repo.urls[short]; present {
		return false
	}
	repo.urls[short] = orig
	return true
}

