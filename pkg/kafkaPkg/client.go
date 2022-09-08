package kafkaPkg

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"prConsumer/config"
)

type client struct {
	KafkaConfig *config.KafkaConfig
	Producer    *kafka.Producer
	Consumers   map[*config.ConsumerConfig]*kafka.Consumer
}

func NewKafkaClient(kafkaConfig *config.KafkaConfig) *client {
	client := &client{
		KafkaConfig: kafkaConfig,
		Producer:    new(kafka.Producer),
		Consumers:   make(map[*config.ConsumerConfig]*kafka.Consumer),
	}
	client.registerProducer()
	client.registerConsumers()
	client.subscribeTopics()
	fmt.Println("Everything is set.")
	return client
}
func (c *client) registerProducer() {
	producerConfigMap := make(kafka.ConfigMap)
	for producerConfigKey, producerConfigValue := range c.KafkaConfig.ProducerConfig {
		producerConfigMap[producerConfigKey] = producerConfigValue
	}
	newProducer, _ := kafka.NewProducer(&producerConfigMap)
	c.Producer = newProducer
}
func (c *client) registerConsumers() {
	for _, consumerConfig := range c.KafkaConfig.ConsumersConfig {
		consumerConfigMap := make(kafka.ConfigMap)
		for configName, configValue := range consumerConfig.Configs {
			consumerConfigMap[configName] = configValue
		}
		newConsumer, _ := kafka.NewConsumer(&consumerConfigMap)
		consumerConfigCopy := consumerConfig
		c.Consumers[&consumerConfigCopy] = newConsumer
	}
}
func (c *client) subscribeTopics() {
	for _, topicConfig := range c.KafkaConfig.TopicsConfig {
		for consumerConfig, consumer := range c.Consumers {
			for _, consumerName := range topicConfig.Consumers {
				if consumerConfig.Name == consumerName {
					_ = consumer.Subscribe(topicConfig.Name, nil)
				}
			}
		}
	}
}
