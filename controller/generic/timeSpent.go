package generic

import (
	"context"
	logger "dct/logger"
	model "dct/model"
	repo "dct/repository"
)

func CreateTimeSpent(ctx context.Context, timeSpent interface{}) error {
	log := logger.GetLoggerFromContext(ctx)
	err := repo.CreateTimeSpent(ctx, timeSpent)
	if err != nil {
		log.Errorf("error creating timeSpent: %v, error: %v", timeSpent, err)
	}
	return err
}

func GetTimeSpent(ctx context.Context, filters map[string]interface{}) (*model.TimeSpent, error) {
	log := logger.GetLoggerFromContext(ctx)
	timeSpent, err := repo.GetTimeSpent(ctx, filters)
	if err != nil {
		log.Errorf("error fetching timeSpent, filters: %v, error: %v", filters, err)
	}
	return timeSpent, err
}

func GetTimeSpentById(ctx context.Context, id int64) (*model.TimeSpent, error) {
	log := logger.GetLoggerFromContext(ctx)
	timeSpent, err := repo.GetTimeSpentById(ctx, id)
	if err != nil {
		log.Errorf("error fetching timeSpent, id: %v, error: %v", id, err)
	}
	return timeSpent, err
}

func GetTimeSpents(ctx context.Context, filters map[string]interface{}) ([]model.TimeSpent, error) {
	log := logger.GetLoggerFromContext(ctx)
	timeSpents, err := repo.GetTimeSpents(ctx, filters)
	if err != nil {
		log.Errorf("error fetching timeSpents, filters: %v, error: %v", filters, err)
	}
	return timeSpents, err
}

func GetAllTimeSpentSum(ctx context.Context, filters map[string]interface{}) (int64, error) {
	log := logger.GetLoggerFromContext(ctx)
	timeSpent, err := repo.GetAllTimeSpentSum(ctx, filters)
	if err != nil {
		log.Errorf("error fetching all time spent sum, filters: %v, error: %v", filters, err)
	}
	return timeSpent, err
}
