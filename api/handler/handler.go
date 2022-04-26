package handler

import (
	"context"
	"dct/logger"
	"net/http"

	apiController "dct/api/controller"
	apiDeserializer "dct/api/deserializer"

	"github.com/gin-gonic/gin"
)

func InsertHandler(ginCtx *gin.Context) {
	log := logger.GetLoggerFromRequest(ginCtx)

	log.Tracef("enter InsertHandler")
	defer log.Tracef("exit InsertHandler")

	reqBody, err := apiDeserializer.InsertDeserializer(ginCtx)
	if err != nil {
		log.Errorf("error parsing requestBody, error: %v", err)
		ginCtx.JSON(http.StatusBadRequest, "bad request")
	}

	log.Infof("requestBody InsertHandler: %+v", reqBody)

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "logger", log)

	respBody, err := apiController.InsertController(ctx, reqBody)
	if err != nil {
		log.Errorf("error in InsertController, error: %v", err)
		ginCtx.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}
	log.Info("responseBody InsertHandler: %v", respBody)
	ginCtx.JSON(http.StatusOK, respBody)
}

func QueryHandler(ginCtx *gin.Context) {

	log := logger.GetLoggerFromRequest(ginCtx)

	log.Tracef("enter QueryHandler")
	defer log.Tracef("exit QueryHandler")

	reqBody, err := apiDeserializer.QueryDeserializer(ginCtx)
	if err != nil {
		log.Errorf("error parsing requestBody, error: %v", err)
		ginCtx.JSON(http.StatusBadRequest, "bad request")
	}

	log.Infof("requestBody QueryHandler: %+v", reqBody)

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "logger", log)

	respBody, err := apiController.QueryController(ctx, reqBody)
	if err != nil {
		log.Errorf("error in QueryController, error: %v", err)
		ginCtx.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}
	log.Info("responseBody QueryHandler: %v", respBody)
	ginCtx.JSON(http.StatusOK, respBody)
}
