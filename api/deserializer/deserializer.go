package deserializer

import (
	"errors"

	apiDtos "dct/utils/dtos/api"

	"github.com/gin-gonic/gin"
)

func InsertDeserializer(ginCtx *gin.Context) (*apiDtos.InsertReqBody, error) {
	var reqBody apiDtos.InsertReqBody
	if err := ginCtx.BindJSON(&reqBody); err != nil {
		return nil, err
	}
	if len(reqBody.Dimensions) == 0 || len(reqBody.Metrics) == 0 {
		return nil, errors.New("not enough data to proceed")
	}
	return &reqBody, nil
}

func QueryDeserializer(ginCtx *gin.Context) (*apiDtos.QueryReqBody, error) {
	var reqBody apiDtos.QueryReqBody
	var err error
	if err = ginCtx.BindJSON(&reqBody); err != nil {
		return nil, err
	}

	if len(reqBody.Dimensions) == 0 {
		return nil, errors.New("not enough data to proceed")
	}
	return &reqBody, nil
}
