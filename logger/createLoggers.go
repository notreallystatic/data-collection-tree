package logger

import (
	"io"
	"os"
	"strings"

	loggerConstant "dct/utils/constants/logger"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CreatePrimaryLogger() *log.Entry {
	return createLoggerToFile(loggerConstant.PrimaryLogFileName, loggerConstant.PrimaryLoggerName)
}

func CreateReqRespLogger() *log.Entry {
	return createLoggerToFile(loggerConstant.ReqRespLogFileName, loggerConstant.ReqRespLoggerName)
}

func createLoggerToFile(logFileName string, loggerName string) *log.Entry {
	hostname := os.Getenv(loggerConstant.HostNameKey)

	i := strings.LastIndex(logFileName, ".")
	logFileName = logFileName[:i] + "_" + hostname + logFileName[i:]

	logWriter := getRotatingLogFileWriter(getLogsDirFilePath() + "/" + logFileName)
	return createLogger(logWriter, loggerName)
}

func getLogsDirFilePath() string {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return curDir + "/logs"
}

func createLogger(outputWriter io.Writer, loggerName string) *log.Entry {
	logger := log.New()
	logger.Out = outputWriter
	// logger.Out = io.MultiWriter(outputWriter, os.Stdout)

	logger.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyMsg:  "message",
			log.FieldKeyFunc: "funcName",
			log.FieldKeyFile: "fileName",
		},
	})

	logger.ReportCaller = true

	loggerEntry := logger.WithFields(
		log.Fields{
			"logger_name":      loggerName,
			"application_name": loggerConstant.ServiceName,
			"pid":              os.Getpid(),
		},
	)

	return loggerEntry
}

func getRotatingLogFileWriter(logFilePath string) io.Writer {
	logger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    10,
		MaxBackups: 1000,
	}

	return io.Writer(logger)
}
