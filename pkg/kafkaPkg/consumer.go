package kafkaPkg

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	uuid "github.com/satori/go.uuid"
	"prConsumer/events"
	"prConsumer/handlerPkg"
	"prConsumer/model"
	"time"
)

type Consumer struct {
	appHandler *handlerPkg.Handler
}

var appConsumer *Consumer

func SetAppConsumer(appHandler *handlerPkg.Handler) {
	appConsumer = new(Consumer)
	appConsumer.appHandler = appHandler
}

func (c *client) StartConsuming() {
	for {
		Consume(c)
	}
}

func Consume(c *client) {
	for consumerConfig, consumer := range c.Consumers {
		consumer := consumer
		consumerConfig := consumerConfig
		go func() {
			event := consumer.Poll(0)
			switch event := event.(type) {
			case *kafka.Message:
				if err := consumerConfig.ConsumeMethod(event); err != nil {
					RunRetry(event, c)
				}
			}
		}()
	}
}
func OrderCreatedConsume(message *kafka.Message) error {
	var orderCreated events.OrderCreated
	if err := json.Unmarshal(message.Value, &orderCreated); err != nil {
		panic("UNMARSHALL ERROR")
	}
	log := &model.Log{
		Id:      uuid.NewV4().String(),
		Message: fmt.Sprintf("Order is created with id:%s, at:%s", orderCreated.Id, orderCreated.CreatedAt),
	}
	if err := appConsumer.appHandler.LogHandler.Log(log); err != nil {
		return err
	}
	return nil

}
func Order4snConsume(message *kafka.Message) error {
	time.Sleep(4 * time.Second)
	return OrderCreatedConsume(message)
}
func Order8snConsume(message *kafka.Message) error {
	time.Sleep(8 * time.Second)
	return OrderCreatedConsume(message)
}

func RunRetry(message *kafka.Message, kafkaClient *client) {
	retryCount := GetRetryCountHeader(message)
	switch *retryCount {
	case 0:
		SetRetryCountHeader(message, "1")
		kafkaClient.Produce(message, kafkaClient.KafkaConfig.TopicsConfig.Order4sn.Name)
	case 1:
		SetRetryCountHeader(message, "2")
		kafkaClient.Produce(message, kafkaClient.KafkaConfig.TopicsConfig.Order8sn.Name)
	case 2:
		kafkaClient.Produce(message, kafkaClient.KafkaConfig.TopicsConfig.DeadLetter.Name)
	}

}
func GetRetryCountHeader(message *kafka.Message) *int {
	retryCount := 0
	for _, header := range message.Headers {
		if header.Key == "retryCount" {
			err := json.Unmarshal(header.Value, &retryCount)
			if err != nil {
				retryCount = 0
			}
		}
	}
	return &retryCount
}
func SetRetryCountHeader(message *kafka.Message, rc string) {
	retryHeader := kafka.Header{
		Key:   "retryCount",
		Value: []byte(rc),
	}
	for index, header := range message.Headers {
		if header.Key == retryHeader.Key {
			message.Headers[index] = retryHeader
			return
		}
	}
	message.Headers = append(message.Headers, retryHeader)
}
