package kafkaPkg

import "github.com/confluentinc/confluent-kafka-go/kafka"

func (c *client) Publish(message *kafka.Message, topicName string) {
	for producerConfig, producer := range c.Producers {
		if producerConfig.Topic == topicName {
			deliveryChan := make(chan kafka.Event)
			message.TopicPartition
			producer.Produce(message)
		}

	}

}
