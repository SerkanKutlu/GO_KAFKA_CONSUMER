package internal

import (
	"prConsumer/model"
	mongoPkg "prConsumer/repository/mongo"
)

type LogService struct {
	MongoService *mongoPkg.MongoService
}

var Logger *LogService

func SetLogService(mongoService *mongoPkg.MongoService) {
	Logger = &LogService{
		mongoService,
	}
}

func (logger *LogService) Log(log *model.Log) error {
	if err := logger.MongoService.Insert(log); err != nil {
		return err
	}
	return nil
}
