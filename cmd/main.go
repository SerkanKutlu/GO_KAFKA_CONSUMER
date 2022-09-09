package main

import (
	"prConsumer/config"
	"prConsumer/handlerPkg"
	"prConsumer/internal"
	"prConsumer/pkg/kafkaPkg"
	mongoPkg "prConsumer/repository/mongo"
)

func main() {
	env := "dev"
	//Getting Configurations
	confManager := config.NewConfigurationManager("../yml", "application", env)
	kafkaConf := confManager.GetKafkaConfiguration()
	mongoConf := confManager.GetMongoConfiguration()

	//Mongo service
	mongoService := mongoPkg.GetMongoService(mongoConf)

	//LogRepository
	mongoPkg.SetLogRepository(mongoService)
	logRepository := mongoPkg.GetLogRepository()

	//Log Service
	internal.SetLogService(logRepository)
	logService := internal.GetLogService()

	//Handler
	handlerPkg.SetHandler(logService)
	appHandler := handlerPkg.GetHandler()

	//Kafka package client
	client := kafkaPkg.NewKafkaClient(kafkaConf)
	kafkaPkg.SetAppConsumer(appHandler)
	//Start consuming
	client.StartConsuming()
}
