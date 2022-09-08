package mongoPkg

import (
	"context"
	"prConsumer/model"
)

func (ms MongoService) Insert(log *model.Log) error {
	_, err := ms.Logs.InsertOne(context.Background(), log)
	if err != nil {
		return err
	}
	return nil
}
