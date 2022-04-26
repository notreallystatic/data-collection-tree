package generic

import (
	"context"
	logger "dct/logger"
	model "dct/model"
	repo "dct/repository"
)

func CreateDevice(ctx context.Context, device interface{}) error {
	log := logger.GetLoggerFromContext(ctx)
	err := repo.CreateDevice(ctx, device)
	if err != nil {
		log.Errorf("error creating device: %v, error: %v", device, err)
	}
	return err
}

func GetDevice(ctx context.Context, filters map[string]interface{}) (*model.Device, error) {
	log := logger.GetLoggerFromContext(ctx)
	device, err := repo.GetDevice(ctx, filters)
	if err != nil {
		log.Errorf("error fetching device, filters: %v, error: %v", filters, err)
	}
	return device, err
}

func GetDeviceById(ctx context.Context, id int64) (*model.Device, error) {
	log := logger.GetLoggerFromContext(ctx)
	device, err := repo.GetDeviceById(ctx, id)
	if err != nil {
		log.Errorf("error fetching device by id: %v, error: %v", id, err)
	}
	return device, err
}

func GetDevices(ctx context.Context, filters map[string]interface{}) ([]model.Device, error) {
	log := logger.GetLoggerFromContext(ctx)
	devices, err := repo.GetDevices(ctx, filters)
	if err != nil {
		log.Errorf("error fetching devices, filters: %v, error: %v", filters, err)
	}
	return devices, err
}
