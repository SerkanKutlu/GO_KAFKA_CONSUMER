package config

import "github.com/confluentinc/confluent-kafka-go/kafka"

type AppConfig struct {
	MongoConfig MongoConfig `yaml:"mongoConfig"`
	KafkaConfig KafkaConfig `yaml:"kafkaConfig"`
}

type MongoConfig struct {
	ConnectionString string            `yaml:"connectionString"`
	Database         string            `yaml:"database"`
	Collection       map[string]string `yaml:"collection"`
}

type KafkaConfig struct {
	ProducerConfig  map[string]string  `yaml:"producerConfig"`
	ConsumersConfig ConsumerConfigList `yaml:"consumersConfig"`
	TopicsConfig    TopicConfigList    `yaml:"topicsConfig"`
}

type ConsumerConfigList struct {
	OrderCreatedLogConsumer ConsumerConfig `yaml:"orderCreatedLogConsumer"`
	OrderCreated4snConsumer ConsumerConfig `yaml:"orderCreated4SnConsumer"`
	OrderCreated8snConsumer ConsumerConfig `yaml:"orderCreated8SnConsumer"`
}

type TopicConfigList struct {
	OrderCreated TopicConfig `yaml:"orderCreated"`
	Order4sn     TopicConfig `yaml:"order4Sn"`
	Order8sn     TopicConfig `yaml:"order8Sn"`
	DeadLetter   TopicConfig `yaml:"deadLetter"`
}
type ConsumerConfig struct {
	Name          string            `yaml:"name"`
	Configs       map[string]string `yaml:"configs"`
	Topic         string            `yaml:"topic"`
	ConsumeMethod func(message *kafka.Message) error
}

type TopicConfig struct {
	Name      string   `yaml:"name"`
	Consumers []string `yaml:"consumers"`
	Producers []string `yaml:"producers"`
}
