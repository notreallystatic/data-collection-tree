package main

import (
	logger "dct/logger"
	databaseService "dct/services/database"
	ginService "dct/services/server"
)

func main() {
	databaseService.Init()
	logger.CreateLoggers()
	ginService.Init()
}
