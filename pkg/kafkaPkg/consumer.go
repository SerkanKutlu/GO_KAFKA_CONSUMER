package kafkaPkg

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	uuid "github.com/satori/go.uuid"
	"prConsumer/events"
	"prConsumer/internal"
	"prConsumer/model"
	"time"
)

func (c *client) StartConsuming() {
	c.SetConsumerMethods()
	for {
		for consumerConfig, consumer := range c.Consumers {
			consumer := consumer
			consumerConfig := consumerConfig
			go func() {
				event := consumer.Poll(0)
				switch event := event.(type) {
				case *kafka.Message:
					fmt.Println("message came")
					if err := consumerConfig.ConsumeMethod(event); err != nil {
						RunRetry(event, c)
					}
				}
			}()
		}
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
	if err := internal.Logger.Log(log); err != nil {
		return err
	}
	return nil

}
func Order4snConsume(message *kafka.Message) error {
	fmt.Println("4sn calisti")
	time.Sleep(4 * time.Second)
	return OrderCreatedConsume(message)
}
func Order8snConsume(message *kafka.Message) error {
	fmt.Println("8sn calisti")
	time.Sleep(8 * time.Second)
	return OrderCreatedConsume(message)

}
func (c *client) SetConsumerMethods() {
	for consumerConfig := range c.Consumers {
		if consumerConfig.Name == "ordercreatedlogconsumer" {
			consumerConfig.ConsumeMethod = OrderCreatedConsume
			continue
		}
		if consumerConfig.Name == "ordercreated4snconsumer" {
			consumerConfig.ConsumeMethod = Order4snConsume
			continue
		}
		if consumerConfig.Name == "ordercreated8snconsumer" {
			consumerConfig.ConsumeMethod = Order8snConsume
			continue
		}
	}
}

func RunRetry(message *kafka.Message, kafkaClient *client) {
	retryCount := GetRetryCountHeader(message)
	switch *retryCount {
	case 0:
		SetRetryCountHeader(message, "1")
		kafkaClient.Produce(message, "order4sn")
	case 1:
		SetRetryCountHeader(message, "2")
		kafkaClient.Produce(message, "order8sn")
	case 2:
		kafkaClient.Produce(message, "deadLetter")
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
