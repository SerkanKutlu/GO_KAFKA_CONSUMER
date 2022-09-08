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
	ProducerConfig  map[string]string         `yaml:"producerConfig"`
	ConsumersConfig map[string]ConsumerConfig `yaml:"consumersConfig"`
	TopicsConfig    map[string]TopicConfig    `yaml:"topicsConfig"`
}

type ConsumerConfig struct {
	Name          string            `yaml:"name"`
	Configs       map[string]string `yaml:"configs"`
	ConsumeMethod func(message *kafka.Message) error
}

type TopicConfig struct {
	Name      string   `yaml:"name"`
	Consumers []string `yaml:"consumers"`
	Producers []string `yaml:"producers"`
}
