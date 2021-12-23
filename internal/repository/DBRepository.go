package repository

import (
	"context"
	_ "github.com/lib/pq"
)

func (repo *URLsRepository) CheckPingDB(ctx context.Context) error {
	if err := repo.DB.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func (repo *URLsRepository) CloseDBRepo() error {
	err := repo.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

