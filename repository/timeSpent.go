package repo

import (
	"context"
	model "dct/model"
	database "dct/services/database"
)

func CreateTimeSpent(ctx context.Context, timeSpent interface{}) error {
	tx := database.GetTransactionInstanceFromContext(ctx)
	if tx == nil {
		tx = database.MariaDbClient
	}
	result := tx.Table("time_spent").Create(timeSpent)
	return result.Error
}

func GetTimeSpent(ctx context.Context, filters map[string]interface{}) (*model.TimeSpent, error) {
	var timeSpent model.TimeSpent
	result := database.MariaDbClient.Table("time_spent").Where(filters).First(&timeSpent)
	if result.Error != nil {
		return nil, result.Error
	}
	return &timeSpent, nil
}

func GetTimeSpentById(ctx context.Context, id int64) (*model.TimeSpent, error) {
	return GetTimeSpent(ctx, map[string]interface{}{
		"id": id,
	})
}

func GetAllTimeSpentSum(ctx context.Context, filters map[string]interface{}) (int64, error) {
	var timeSpent *int64
	result := database.MariaDbClient.Table("time_spent").Select("sum(count)").Where(filters).Scan(&timeSpent)
	if timeSpent == nil {
		return 0, result.Error
	}
	return *timeSpent, result.Error
}

func GetTimeSpents(ctx context.Context, filters map[string]interface{}) ([]model.TimeSpent, error) {
	var timeSpents []model.TimeSpent
	result := database.MariaDbClient.Table("time_spent").Where(filters).Find(&timeSpents)
	if result.Error != nil {
		return nil, result.Error
	}
	return timeSpents, nil
}
