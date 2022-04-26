package generic

import (
	"context"
	"dct/logger"
	model "dct/model"
	repo "dct/repository"
)

func CreateCountry(ctx context.Context, country interface{}) error {
	log := logger.GetLoggerFromContext(ctx)
	err := repo.CreateCountry(ctx, country)
	if err != nil {
		log.Errorf("error creating country: %v, error: %v", country, err)
	}
	return err
}

func GetCountry(ctx context.Context, filters map[string]interface{}) (*model.Country, error) {
	log := logger.GetLoggerFromContext(ctx)
	country, err := repo.GetCountry(ctx, filters)
	if err != nil {
		log.Errorf("error fetching country, filters: %+v, error: %v", filters, err)
	}
	return country, err
}

func GetCountryById(ctx context.Context, id int64) (*model.Country, error) {
	log := logger.GetLoggerFromContext(ctx)
	country, err := repo.GetCountryById(ctx, id)
	if err != nil {
		log.Errorf("error fetching country, id: %v, error: %v", id, err)
	}
	return country, err
}

func GetCountries(ctx context.Context, filters map[string]interface{}) ([]model.Country, error) {
	log := logger.GetLoggerFromContext(ctx)
	countries, err := repo.GetCountries(ctx, filters)
	if err != nil {
		log.Errorf("error fetching countries, filters: %+v, error: %v", filters, err)
	}
	return countries, err
}
