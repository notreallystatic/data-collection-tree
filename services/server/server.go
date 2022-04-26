package server

import (
	logger "dct/logger"
	"dct/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

func Init() {
	ginRouter := gin.Default()
	ginRouter.Use(gin.CustomRecovery(customRecovery))
	ginRouter.Use(setLoggerInRequest)
	router.Init(&ginRouter.RouterGroup)
	ginRouter.GET("/health", func(ctx *gin.Context) {
		fmt.Println("health check api")
		ctx.JSON(http.StatusOK, "health well")
	})
	ginRouter.Run()
}

func customRecovery(ctx *gin.Context, err interface{}) {
	log := logger.GetLoggerFromRequest(ctx)
	log.Errorf("exception in server: %v", err)
	ctx.JSON(http.StatusInternalServerError, "an error occurred.")
}

func setLoggerInRequest(ctx *gin.Context) {
	logger := logger.Primary.WithField("request_id", uuid.New())
	ctx.Set("logger", logger)
}
