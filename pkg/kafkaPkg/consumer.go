package kafkaPkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	uuid "github.com/satori/go.uuid"
	"prConsumer/config"
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
		event := consumer.Poll(0)
		switch event := event.(type) {
		case *kafka.Message:
			fmt.Println("GELEN MESAJ BU CONSUMERDE:")
			fmt.Println(consumerConfig.Name)
			addToBatch(event, consumerConfig)
			if c.isBatchReady(consumerConfig) {
				fmt.Println("batch is ready consuming")
				for _, message := range *consumerConfig.Batch {
					fmt.Println("message id")
					fmt.Println(&message)
					if err := consumerConfig.ConsumeMethod(&message); err != nil {
						go RunRetry(&message, c)
					}
				}
				resetBatch(consumerConfig)
			} else {
				fmt.Println("Batch is not ready:")
				fmt.Println(consumerConfig.BatchSize)
				fmt.Println(consumerConfig.MessageCount)
			}

		}
	}
}
func addToBatch(message *kafka.Message, consumerConfig *config.ConsumerConfig) {
	if consumerConfig.Batch == nil {
		consumerConfig.Batch = new([]kafka.Message)
	}
	*consumerConfig.Batch = append(*consumerConfig.Batch, *message)
	consumerConfig.MessageCount++
	consumerConfig.BatchSize += len(message.Value)
}
func (c *client) isBatchReady(consumerConfig *config.ConsumerConfig) bool {

	batchSize := c.KafkaConfig.ConsumerBatchSettings.BatchSize
	messageCount := c.KafkaConfig.ConsumerBatchSettings.MessageCount
	return consumerConfig.BatchSize >= batchSize || consumerConfig.MessageCount >= messageCount

}

func resetBatch(consumerConfig *config.ConsumerConfig) {
	consumerConfig.Batch = nil
	consumerConfig.Batch = new([]kafka.Message)
	consumerConfig.BatchSize = 0
	consumerConfig.MessageCount = 0
}
func OrderCreatedConsume(message *kafka.Message) error {
	fmt.Println("OrderCreatedConsume consumer worked")
	return errors.New("xxx")
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
	fmt.Println("4sn consumer worked")
	time.Sleep(4 * time.Second)
	return OrderCreatedConsume(message)
}
func Order8snConsume(message *kafka.Message) error {
	fmt.Println("8sn consumer worked")
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
