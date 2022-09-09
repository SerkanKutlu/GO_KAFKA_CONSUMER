package mongoPkg

import (
	"context"
	"prConsumer/model"
)

type LogRepository struct {
	mongoService *MongoService
}

var logRepository *LogRepository

func SetLogRepository(ms *MongoService) {
	logRepository = new(LogRepository)
	logRepository.mongoService = ms
}
func GetLogRepository() *LogRepository {
	return logRepository
}
func (lr *LogRepository) Insert(log *model.Log) error {
	_, err := lr.mongoService.Logs.InsertOne(context.Background(), log)
	if err != nil {
		return err
	}
	return nil
}
