package kafkaPkg

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"prConsumer/config"
)

type client struct {
	KafkaConfig *config.KafkaConfig
	Producer    *kafka.Producer
	Consumers   map[*config.ConsumerConfig]*kafka.Consumer
}
type ConsumeMethod func(message *kafka.Message) error

func NewKafkaClient(kafkaConfig *config.KafkaConfig) *client {
	client := &client{
		KafkaConfig: kafkaConfig,
		Producer:    new(kafka.Producer),
		Consumers:   make(map[*config.ConsumerConfig]*kafka.Consumer),
	}
	client.registerProducer()
	client.registerConsumers()
	client.subscribeTopics()
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
	//OrderCreated4sn Consumer Register
	consumerConfig := &c.KafkaConfig.ConsumersConfig.OrderCreated4snConsumer
	kafkaConsumer := createConsumer(consumerConfig, Order4snConsume)
	c.Consumers[consumerConfig] = kafkaConsumer

	//OrderCreated8sn Consumer Register
	consumerConfig = &c.KafkaConfig.ConsumersConfig.OrderCreated8snConsumer
	kafkaConsumer = createConsumer(consumerConfig, Order8snConsume)
	c.Consumers[consumerConfig] = kafkaConsumer

	//OrderCreatedLog Consumer Register
	consumerConfig = &c.KafkaConfig.ConsumersConfig.OrderCreatedLogConsumer
	kafkaConsumer = createConsumer(consumerConfig, OrderCreatedConsume)
	c.Consumers[consumerConfig] = kafkaConsumer

}
func (c *client) subscribeTopics() {
	for consumerConfig, kafkaConsumer := range c.Consumers {
		_ = kafkaConsumer.Subscribe(consumerConfig.Topic, nil)
	}
}
func createConsumer(consumerConfig *config.ConsumerConfig, consumeMethod ConsumeMethod) *kafka.Consumer {
	consumerConfigMap := make(kafka.ConfigMap)
	for key, value := range consumerConfig.Configs {
		consumerConfigMap[key] = value
	}
	consumerConfig.ConsumeMethod = consumeMethod
	kafkaConsumer, _ := kafka.NewConsumer(&consumerConfigMap)
	return kafkaConsumer
}
