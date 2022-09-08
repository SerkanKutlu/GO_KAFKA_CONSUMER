package config

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
	ProducersConfig map[string]ProducerConfig `yaml:"producersConfig"`
	ConsumersConfig map[string]ConsumerConfig `yaml:"consumersConfig"`
	TopicsConfig    map[string]TopicConfig    `yaml:"topicsConfig"`
}

type ConsumerConfig struct {
	Name    string            `yaml:"name"`
	Configs map[string]string `yaml:"configs"`
}

type ProducerConfig struct {
	Name    string            `yaml:"name"`
	Topic   string            `yaml:"topic"`
	Configs map[string]string `yaml:"configs"`
}
type TopicConfig struct {
	Name      string   `yaml:"name"`
	Consumers []string `yaml:"consumers"`
	Producers []string `yaml:"producers"`
}
