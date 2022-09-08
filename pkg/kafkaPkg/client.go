package kafkaPkg

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"prConsumer/config"
)

type client struct {
	KafkaConfig *config.KafkaConfig
	Producers   map[*config.ProducerConfig]*kafka.Producer
	Consumers   map[*config.ConsumerConfig]*kafka.Consumer
}

func NewKafkaClient(kafkaConfig *config.KafkaConfig) {
	client := &client{
		KafkaConfig: kafkaConfig,
		Producers:   make(map[*config.ProducerConfig]*kafka.Producer),
		Consumers:   make(map[*config.ConsumerConfig]*kafka.Consumer),
	}
	client.registerProducers()
	client.registerConsumers()
	client.subscribeTopics()
	fmt.Println("Everything is set.")
}
func (c *client) registerProducers() {

	for _, producerConfig := range c.KafkaConfig.ProducersConfig {
		producerConfigMap := make(kafka.ConfigMap)
		for configName, configValue := range producerConfig.Configs {
			producerConfigMap[configName] = configValue
		}
		newProducer, _ := kafka.NewProducer(&producerConfigMap)
		c.Producers[&producerConfig] = newProducer
	}
}
func (c *client) registerConsumers() {
	for _, consumerConfig := range c.KafkaConfig.ConsumersConfig {
		consumerConfigMap := make(kafka.ConfigMap)
		for configName, configValue := range consumerConfig.Configs {
			consumerConfigMap[configName] = configValue
		}
		newConsumer, _ := kafka.NewConsumer(&consumerConfigMap)
		c.Consumers[&consumerConfig] = newConsumer
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
