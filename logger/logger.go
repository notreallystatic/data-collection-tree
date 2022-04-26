package logger

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var Primary *log.Entry
var ReqResp *log.Entry

func CreateLoggers() {
	Primary = CreatePrimaryLogger()
	Primary.Logger.SetLevel(log.TraceLevel)
	Primary.Logger.SetNoLock()
	ReqResp = CreateReqRespLogger()
}

type WriteFunc func([]byte) (int, error)

func (fn WriteFunc) Write(data []byte) (int, error) {
	return fn(data)
}

func GetLoggerFromRequest(ctx *gin.Context) *logrus.Entry {
	if logger, exists := ctx.Get("logger"); exists {
		if logger, ok := logger.(*logrus.Entry); ok {
			return logger
		}
	}
	return Primary
}

func GetLoggerFromContext(ctx context.Context) *logrus.Entry {
	if val := ctx.Value("logger"); val != nil {
		if logger, ok := val.(*logrus.Entry); ok {
			return logger
		}
	}
	return Primary
}
