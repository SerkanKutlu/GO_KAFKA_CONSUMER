package mongoPkg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"prConsumer/config"
	"time"
)

type MongoService struct {
	Logs *mongo.Collection
}

func GetMongoService(mongoConfig *config.MongoConfig) *MongoService {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConfig.ConnectionString))
	if err != nil {
		panic(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	database := client.Database(mongoConfig.Database)
	logs := database.Collection(mongoConfig.Collection["log"])
	service := MongoService{
		logs,
	}
	return &service
}
