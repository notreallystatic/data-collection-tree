package repo

import (
	"context"
	model "dct/model"
	database "dct/services/database"
)

func CreateDevice(ctx context.Context, country interface{}) error {
	tx := database.GetTransactionInstanceFromContext(ctx)
	if tx == nil {
		tx = database.MariaDbClient
	}
	result := tx.Table("device").Create(country)
	return result.Error
}

func GetDevice(ctx context.Context, filters map[string]interface{}) (*model.Device, error) {
	var device model.Device
	result := database.MariaDbClient.Table("device").Where(filters).First(&device)
	if result.Error != nil {
		return nil, result.Error
	}
	return &device, nil
}

func GetDeviceById(ctx context.Context, id int64) (*model.Device, error) {
	return GetDevice(ctx, map[string]interface{}{
		"id": id,
	})
}

func GetDevices(ctx context.Context, filters map[string]interface{}) ([]model.Device, error) {
	var devices []model.Device
	result := database.MariaDbClient.Table("device").Where(filters).Find(&devices)
	if result.Error != nil {
		return nil, result.Error
	}
	return devices, nil
}
