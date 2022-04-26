package generic

import (
	"context"
	logger "dct/logger"
	model "dct/model"
	repo "dct/repository"
)

func CreateWebRequest(ctx context.Context, webRequest interface{}) error {
	log := logger.GetLoggerFromContext(ctx)
	err := repo.CreateWebRequest(ctx, webRequest)
	if err != nil {
		log.Errorf("error creating webRequest: %v, error: %v", webRequest, err)
	}
	return err
}

func GetWebRequest(ctx context.Context, filters map[string]interface{}) (*model.WebRequest, error) {
	log := logger.GetLoggerFromContext(ctx)
	webRequest, err := repo.GetWebRequest(ctx, filters)
	if err != nil {
		log.Errorf("error fetching webRequest: %v, error: %v", filters, err)
	}
	return webRequest, err
}

func GetWebRequestById(ctx context.Context, id int64) (*model.WebRequest, error) {
	log := logger.GetLoggerFromContext(ctx)
	webRequest, err := repo.GetWebRequestById(ctx, id)
	if err != nil {
		log.Errorf("error fetching webRequest by id: %v, error: %v", id, err)
	}
	return webRequest, err
}

func GetWebRequests(ctx context.Context, filters map[string]interface{}) ([]model.WebRequest, error) {
	log := logger.GetLoggerFromContext(ctx)
	webRequests, err := repo.GetWebRequests(ctx, filters)
	if err != nil {
		log.Errorf("error fetching webRequests: %v, error: %v", filters, err)
	}
	return webRequests, err
}

func GetAllWebRequestsSum(ctx context.Context, filters map[string]interface{}) (int64, error) {
	log := logger.GetLoggerFromContext(ctx)
	webRequests, err := repo.GetAllWebRequestsSum(ctx, filters)
	if err != nil {
		log.Errorf("error fetching all web requests sum, filters: %v, error: %v", filters, err)
	}
	return webRequests, err
}
