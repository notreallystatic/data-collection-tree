package router

import (
	apiHandlers "dct/api/handler"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	{
		v1Group.POST("/insert", apiHandlers.InsertHandler)
		v1Group.GET("/query", apiHandlers.QueryHandler)
	}
}
