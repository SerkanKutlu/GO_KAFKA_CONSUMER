package internal

import (
	mongoPkg "prConsumer/repository/mongo"
)

type LogService struct {
	LogRepository *mongoPkg.LogRepository
}

var Logger *LogService

func SetLogService(logRepository *mongoPkg.LogRepository) {
	Logger = &LogService{
		logRepository,
	}
}

func GetLogService() *LogService {
	return Logger
}
