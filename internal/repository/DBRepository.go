package repository

import (
	"context"
	"time"
)

func (repo *URLsRepository) CheckPingDB(ctx context.Context) error {
	timeoutDur := time.Second * 1
	ctx2, cancel := context.WithTimeout(ctx, timeoutDur)
	defer cancel()

	if err := repo.DB.PingContext(ctx2); err != nil {
		return err
	}
	return nil
}


