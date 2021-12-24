package services

import "context"

func (service *Service) CreateTableDB(ctx context.Context, name string) error {
	err := service.CreateTableDBRepo(ctx, name)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) InsertUserDB(ctx context.Context, tabname string, userID string, shortURL string, origURL string) error {
	err := service.InsertUserDBRepo(ctx, tabname, userID, shortURL, origURL)
	if err != nil {
		return err
	}
	return nil
}
