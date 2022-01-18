package repository_db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/key_gen"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

type DBRepository struct {
	DB *sql.DB
}

func (s *DBRepository) New(cfg *config.Config) {
	db, err := sql.Open("postgres", cfg.DatabaseDSN)
	if err != nil {
		log.Fatal("error connecting to DB:", err)
	}

	s.DB = db

	ctx, stop := context.WithTimeout(context.Background(), 1*time.Second)
	defer stop()

	result, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users(id SERIAL, user_id VARCHAR, short_url VARCHAR NOT NULL, orig_url VARCHAR NOT NULL, correlation_id VARCHAR, is_deleted VARCHAR(10) DEFAULT '', PRIMARY KEY (id), UNIQUE (orig_url));`)
	if err != nil {
		log.Fatal("error create table in DB", err)
	}

	log.Println(result)

}

func (s *DBRepository) CheckPingRepo(ctx context.Context) error {
	log.Println("IN DB:")
	if err := s.DB.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func (s *DBRepository) InsertUserRepo(ctx context.Context, userID string, shortURL string, origURL string) (string, error) {
	log.Println("IN DB:")
	db := s.DB

	if db == nil {
		log.Println("You haven`t open the database connection")
		return "", errors.New("you haven`t open the database connection")
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	var result string

	log.Println("IN InsertUserRepo short url =", shortURL)
	_, err := db.ExecContext(ctx, `INSERT INTO users(user_id, short_url, orig_url) VALUES ($1, $2, $3);`, userID, shortURL, origURL)

	if err != nil {
		if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
			notOk := db.QueryRowContext(ctx, `SELECT short_url FROM users WHERE orig_url=$1;`, origURL).Scan(&result)
			if notOk != nil {
				return "", notOk
			}
			return result, err
		}
	}

	return shortURL, nil
}

func (s *DBRepository) GetOrigByShortRepo(ctx context.Context, shortURL string) (string, error) {
	log.Println("IN DB:")
	var result string
	var deleted string
	if s.DB == nil {
		log.Println("You haven`t open the database connection")
		return "", errors.New("you haven`t open the database connection")
	}

	db := s.DB

	ctx2, stop := context.WithTimeout(ctx, 1*time.Second)
	defer stop()

	log.Println("In GetOrigByShortRepo: short url =", shortURL)

	err := db.QueryRowContext(ctx2, `SELECT orig_url, is_deleted FROM users WHERE short_url=$1;`, shortURL).Scan(&result, &deleted)
	if err != nil {
		return "", err
	}

	log.Printf("result=%s, deleted=%s", result, deleted)

	if deleted == "deleted" {
		result = deleted
	}

	return result, nil
}

func (s *DBRepository) GetURLsListRepo(ctx context.Context, uuid string) ([]models.URLs, error) {
	log.Println("IN DB:")
	var user []models.URLs

	db := s.DB

	ctx2, stop := context.WithTimeout(ctx, 1*time.Second)
	defer stop()

	rows, err := db.QueryContext(ctx2, `SELECT user_id, short_url, orig_url FROM users WHERE user_id=$1`, uuid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var v models.URLs
		err = rows.Scan(&v.UUID, &v.ShortURL, &v.OriginalURL)
		if err != nil {
			return nil, err
		}
		user = append(user, v)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *DBRepository) BanchApiRepo(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error) {
	log.Println("BanchApiRepo IN DB:")
	result := make([]models.BatchResult, 0)

	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users(user_id, short_url, orig_url, correlation_id) 
												VALUES($1, $2, $3, $4)`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for _, v := range in {
		short := key_gen.GenerateShortLink(v.OriginalURL, v.CorrelationID)
		if _, err = stmt.ExecContext(ctx, uuid, short, v.OriginalURL, v.CorrelationID); err != nil {
			log.Println("error insert in db!")
			return nil, err
		}
		result = append(result, models.BatchResult{CorrelationID: v.CorrelationID, ShortURL: shortBase + "/" + short})
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}

func fanIn(inputChs ...chan string) chan string {
	outCh := make(chan string)
	go func() {
		wg := &sync.WaitGroup{}

		for _, inputCh := range inputChs {
			wg.Add(1)
			go func(inputCh chan string) {
				defer wg.Done()
				for item := range inputCh {
					outCh <- item
				}
			}(inputCh)
		}
		wg.Wait()
		close(outCh)
	}()
	return outCh
}

func (s *DBRepository) UpdateURLsRepo(ctx context.Context, uuid string, shortBases []string) error {
	db := s.DB
	if db == nil {
		return errors.New("You haven`t opened the database connection")
	}

	if len(shortBases) == 0 {
		return errors.New("request is empty...")
	}

	n := len(shortBases)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if len(shortBases) == 0 {
		return errors.New("the list of URLs is empty")
	}

	inputChs := make([]chan string, 0, n)

	go func() {
		for _, item := range shortBases {
			itemsCh := make(chan string)
			itemsCh <- item
			inputChs = append(inputChs, itemsCh)
		}

	}()

	stmt, err := tx.PrepareContext(ctx, `UPDATE users SET is_deleted=$1 WHERE user_id=$2 AND short_url=$3`)
	if err != nil {
		log.Println("PrepareContext error...")
		return err
	}

	for item := range fanIn(inputChs...) {
		log.Println("UUID=%s, SHORTURL=%s", uuid, item)
		if _, err = stmt.ExecContext(ctx, "deleted", uuid, item); err != nil {
			if err = tx.Rollback(); err != nil {
				log.Println("Rollback error...")
				return err
			}
			log.Println("ExecContext error...")
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		log.Println("Commit error...")
		return err
	}

	return nil
}
