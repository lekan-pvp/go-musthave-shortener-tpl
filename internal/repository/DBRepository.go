package repository

import (
	"context"
	"fmt"
	"github.com/lib/pq"
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

func (repo *Repository) CreateTableDBRepo(ctx context.Context, tableName string) error {
	db := repo.DB

	ctx2, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	tblname := pq.QuoteIdentifier(tableName)
	_, err := db.ExecContext(ctx2, fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (user_id PRIMARY KEY, short_url VARCHAR(50) NOT NULL, orig_url VARCHAR(50) NOT NULL);", tblname))

	if err != nil {
		log.Println("in CreateTableDBRepo:", err)
		return err
	}
	return nil
}

func (repo *Repository) InsertUserDBRepo(ctx context.Context, tabname string, userID string, shortURL string, origURL string) error {
	db := repo.DB

	ctx2, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	tblname := pq.QuoteIdentifier(tabname)
	id := pq.QuoteIdentifier(userID)
	short := pq.QuoteIdentifier(shortURL)
	orig := pq.QuoteIdentifier(origURL)

	_, err := db.ExecContext(ctx2, fmt.Sprintf("INSERT INTO %s(user_id, short_url, orig_url) VALUES (%s, %s, %s);", tblname, id, short, orig))
	if err != nil {
		log.Println("in InsertUser:", err)
		return err
	}
	return nil
}