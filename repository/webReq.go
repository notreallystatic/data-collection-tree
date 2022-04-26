package repo

import (
	"context"
	model "dct/model"
	database "dct/services/database"
)

func CreateWebRequest(ctx context.Context, webRequest interface{}) error {
	tx := database.GetTransactionInstanceFromContext(ctx)
	if tx == nil {
		tx = database.MariaDbClient
	}
	result := tx.Table("web_request").Create(webRequest)
	return result.Error
}

func GetWebRequest(ctx context.Context, filters map[string]interface{}) (*model.WebRequest, error) {
	var webRequest model.WebRequest
	result := database.MariaDbClient.Table("web_request").Where(filters).First(&webRequest)
	if result.Error != nil {
		return nil, result.Error
	}
	return &webRequest, nil
}

func GetAllWebRequestsSum(ctx context.Context, filters map[string]interface{}) (int64, error) {
	var webRequests *int64
	result := database.MariaDbClient.Table("web_request").Select("sum(count)").Where(filters).Scan(&webRequests)
	if webRequests == nil {
		return 0, result.Error
	}
	return *webRequests, result.Error
}

func GetWebRequestById(ctx context.Context, id int64) (*model.WebRequest, error) {
	return GetWebRequest(ctx, map[string]interface{}{
		"id": id,
	})
}

func GetWebRequests(ctx context.Context, filters map[string]interface{}) ([]model.WebRequest, error) {
	var webRequests []model.WebRequest
	result := database.MariaDbClient.Table("web_request").Where(filters).Find(&webRequests)
	if result.Error != nil {
		return nil, result.Error
	}
	return webRequests, nil
}
