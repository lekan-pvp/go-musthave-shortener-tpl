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


