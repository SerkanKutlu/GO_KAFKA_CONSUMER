package main

import (
	"prConsumer/config"
	"prConsumer/internal"
	"prConsumer/pkg/kafkaPkg"
	mongoPkg "prConsumer/repository/mongo"
)

func main() {
	env := "dev"
	confManager := config.NewConfigurationManager("../yml", "application", env)
	kafkaConf := confManager.GetKafkaConfiguration()
	mongoConf := confManager.GetMongoConfiguration()
	mongoService := mongoPkg.GetMongoService(mongoConf)
	internal.SetLogService(mongoService)

	client := kafkaPkg.NewKafkaClient(kafkaConf)
	client.StartConsuming()
}
