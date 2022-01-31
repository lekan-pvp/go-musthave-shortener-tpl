package repositorydb

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/keygen"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
	"github.com/lib/pq"
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

	result, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users(id SERIAL, user_id VARCHAR, short_url VARCHAR NOT NULL, orig_url VARCHAR NOT NULL, correlation_id VARCHAR, is_deleted BOOLEAN DEFAULT FALSE, PRIMARY KEY (id), UNIQUE (orig_url));`)
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

func (s *DBRepository) GetOrigByShortRepo(ctx context.Context, shortURL string) (*models.OriginLink, error) {
	log.Println("IN DB:")
	result := models.OriginLink{}
	if s.DB == nil {
		log.Println("You haven`t open the database connection")
		return nil, errors.New("you haven`t open the database connection")
	}

	db := s.DB

	log.Println("In GetOrigByShortRepo: short url =", shortURL)

	err := db.QueryRowContext(ctx, `SELECT orig_url, is_deleted FROM users WHERE short_url=$1;`, shortURL).Scan(&result.Link, &result.Deleted)
	if err != nil {
		return nil, err
	}

	log.Printf("result=%s, deleted=%t", result.Link, result.IsDeleted())

	return &result, nil
}

//GetURLsListRepo...
func (s *DBRepository) GetURLsListRepo(ctx context.Context, uuid string) ([]models.URLs, error) {
	log.Println("IN DB:")
	var user []models.URLs

	db := s.DB

	rows, err := db.QueryContext(ctx, `SELECT user_id, short_url, orig_url FROM users WHERE user_id=$1`, uuid)
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

func (s *DBRepository) BanchAPIRepo(ctx context.Context, uuid string, in []models.BatchIn, shortBase string) ([]models.BatchResult, error) {
	log.Println("BanchAPIRepo IN DB:")
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
		short := keygen.GenerateShortLink(v.OriginalURL, v.CorrelationID)
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

func fanOut(input []string, n int) []chan string {
	chs := make([]chan string, 0, n)
	for i, val := range input {
		ch := make(chan string, 1)
		ch <- val
		chs = append(chs, ch)
		close(chs[i])
	}
	return chs
}

func newWorker(ctx context.Context, stmt *sql.Stmt, tx *sql.Tx, inputCh <-chan string) error {
		defer func() {
			select {
			case <-ctx.Done():
				log.Println("aborting dalete:", ctx.Err())
				}
		}()
		for id := range inputCh {
			if _, err := stmt.ExecContext(ctx, id); err != nil {
				if err = tx.Rollback(); err != nil {
						return err
					}
				return err
			}
		}
		return nil
}

func (s *DBRepository) UpdateURLsRepo(ctx context.Context, shortBases []string) error {
	n := len(shortBases)

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	if len(shortBases) == 0 {
		return errors.New("the list of URLs is empty")
	}

	fanOutChs := fanOut(shortBases, n)

	stmt, err := tx.PrepareContext(ctx, `UPDATE users SET is_deleted=TRUE WHERE short_url=$1`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	wg := sync.WaitGroup{}
	errCh := make(chan error)
	jobs := make(chan string, n)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			err := newWorker(ctx, stmt, tx, jobs)
			errCh <- err
			defer wg.Done()
		}()
	}
	wg.Wait()

	for _, item := range fanOutChs {
		jobs <- <-item
	}
	close(jobs)

	if err = <-errCh; err != nil {
		log.Println(err)
		return err
	}
	close(errCh)

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
