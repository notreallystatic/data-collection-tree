package controller

import (
	"context"
	"errors"

	genericController "dct/controller/generic"
	model "dct/model"
	database "dct/services/database"
	apiDtos "dct/utils/dtos/api"
)

func InsertController(ctx context.Context, reqBody *apiDtos.InsertReqBody) (interface{}, error) {

	tx := database.MariaDbClient.Begin()
	ctx = context.WithValue(ctx, "dbTxn", tx)
	defer tx.Rollback()

	var err error
	var deviceInfo *model.Device
	var countryInfo *model.Country

	for _, dim := range reqBody.Dimensions {
		switch dim.Key {
		case "device":
			deviceInfo, err = genericController.GetDevice(ctx, map[string]interface{}{
				"name": dim.Val,
			})
			if err != nil {
				return nil, err
			}
		case "country":
			countryInfo, err = genericController.GetCountry(ctx, map[string]interface{}{
				"keyword": dim.Val,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	if deviceInfo == nil || countryInfo == nil {
		err = errors.New("device or country data is nil")
		return nil, err
	}

	for _, metric := range reqBody.Metrics {
		switch metric.Key {
		case "webreq":
			newWebReq := model.WebRequest{
				CountryId: countryInfo.Id,
				DeviceId:  deviceInfo.Id,
				Count:     metric.Val,
			}
			err = genericController.CreateWebRequest(ctx, &newWebReq)
			if err != nil {
				return nil, err
			}
		case "timespent":
			newTimeSpent := model.TimeSpent{
				CountryId: countryInfo.Id,
				DeviceId:  deviceInfo.Id,
				Count:     metric.Val,
			}
			err = genericController.CreateTimeSpent(ctx, &newTimeSpent)
			if err != nil {
				return nil, err
			}
		}
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	message := "data saved successfully!"
	return message, nil
}

func QueryController(ctx context.Context, reqBody *apiDtos.QueryReqBody) (*apiDtos.QueryRespBody, error) {

	tx := database.MariaDbClient.Begin()
	ctx = context.WithValue(ctx, "dbTxn", tx)

	var respBody apiDtos.QueryRespBody
	var webRequest int64
	var timeSpent int64
	var err error

	for _, dim := range reqBody.Dimensions {
		filters := make(map[string]interface{})
		if dim.Key == "country" {
			country, err := genericController.GetCountry(ctx, map[string]interface{}{
				"keyword": dim.Val,
			})
			if err != nil {
				return nil, err
			}
			filters["country_id"] = country.Id
		} else if dim.Key == "device" {
			device, err := genericController.GetDevice(ctx, map[string]interface{}{
				"name": dim.Val,
			})
			if err != nil {
				return nil, err
			}
			filters["device_id"] = device.Id
		}
		webRequest, err = genericController.GetAllWebRequestsSum(ctx, filters)
		if err != nil {
			return nil, err
		}
		timeSpent, err = genericController.GetAllTimeSpentSum(ctx, filters)
		if err != nil {
			return nil, err
		}
		respBody.Dimensions = append(respBody.Dimensions, dim)
		respBody.Metrics = append(respBody.Metrics, apiDtos.Metric{
			Key: "webreq",
			Val: webRequest,
		})
		respBody.Metrics = append(respBody.Metrics, apiDtos.Metric{
			Key: "timespent",
			Val: timeSpent,
		})
	}
	return &respBody, nil
}
