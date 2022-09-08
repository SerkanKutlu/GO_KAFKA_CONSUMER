package main

import (
	"prConsumer/config"
	"prConsumer/pkg/kafkaPkg"
)

func main() {
	env := "dev"
	confManager := config.NewConfigurationManager("../yml", "application", env)
	kafkaConf := confManager.GetKafkaConfiguration()
	kafkaPkg.NewKafkaClient(kafkaConf)
}
