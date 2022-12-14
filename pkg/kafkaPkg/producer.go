package kafkaPkg

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (c *client) Produce(message *kafka.Message, topicName string) {
	deliveryChan := make(chan kafka.Event, 10000)
	message.TopicPartition.Topic = &topicName
	_ = c.Producer.Produce(message, deliveryChan)
}
