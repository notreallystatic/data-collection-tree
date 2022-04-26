package repo

import (
	"context"
	model "dct/model"
	database "dct/services/database"
)

func CreateCountry(ctx context.Context, country interface{}) error {
	tx := database.GetTransactionInstanceFromContext(ctx)
	if tx == nil {
		tx = database.MariaDbClient
	}
	result := tx.Table("country").Create(country)
	return result.Error
}

func GetCountry(ctx context.Context, filters map[string]interface{}) (*model.Country, error) {
	var country model.Country
	result := database.MariaDbClient.Table("country").Where(filters).First(&country)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}

func GetCountryById(ctx context.Context, id int64) (*model.Country, error) {
	return GetCountry(ctx, map[string]interface{}{
		"id": id,
	})
}

func GetCountries(ctx context.Context, filters map[string]interface{}) ([]model.Country, error) {
	var countries []model.Country
	result := database.MariaDbClient.Table("country").Where(filters).Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return countries, nil
}
