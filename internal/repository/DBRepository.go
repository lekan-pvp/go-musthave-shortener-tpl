package repository

import "context"

func (repo *URLsRepository) CheckPingDB(ctx context.Context) error {
	if err := repo.DB.PingContext(ctx); err != nil {
		return err
	}
	return nil
}


