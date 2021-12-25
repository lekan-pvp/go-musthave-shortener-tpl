package repository

import (
	"context"
	"errors"
	"github.com/go-musthave-shortener-tpl/internal/models"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func (repo *Repository) CheckPingDB(ctx context.Context) error {
	if err := repo.DB.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) CloseDBRepo() error {
	err := repo.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) InsertUserDBRepo(ctx context.Context, userID string, shortURL string, origURL string) error {
	if repo.DB == nil {
		log.Println("You haven`t open the database connection")
		return errors.New("you haven`t open the database connection")
	}

	db := repo.DB

	ctx2, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	log.Println("IN InsertUserDBRepo short url =", shortURL)
	_, err := db.ExecContext(ctx2, `INSERT INTO users(user_id, short_url, orig_url) VALUES ($1, $2, $3);`, userID, shortURL, origURL)
	if err != nil {
		log.Println("in InsertUser:", err)
		return err
	}
	return nil
}

func (repo *Repository) GetOrigByShortDBRepo(ctx context.Context, shortURL string) (string, error)  {
	var result string
	if repo.DB == nil {
		log.Println("You haven`t open the database connection")
		return "", errors.New("you haven`t open the database connection")
	}

	db := repo.DB

	ctx2, stop := context.WithTimeout(ctx, 1*time.Second)
	defer stop()

	log.Println("In GetOrigByShortDBRepo: short url =", shortURL)

	err := db.QueryRowContext(ctx2, `SELECT orig_url FROM users WHERE short_url=$1;`, shortURL).Scan(&result)
	if err != nil {
		return "", err
	}

	log.Println("ORIG URL=", result)

	return result, nil
}

func (repo *Repository) GetURLsListDBRepo(ctx context.Context, uuid string) ([]models.URLs, error) {
	var user []models.URLs

	db := repo.DB

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