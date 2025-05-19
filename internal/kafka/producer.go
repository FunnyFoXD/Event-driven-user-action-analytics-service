package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"strconv"
	"time"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(broker, topic string) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return &Producer{writer: writer}
}

func (p *Producer) SendMessage(ctx context.Context, key string, value []byte) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: value,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

func createTopic(broker, topic string) {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		panic("failed to get controller: " + err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic("failed to get controller: " + err.Error())
	}
	conn.Close()

	conn, err = kafka.Dial("tcp", controller.Host+":"+strconv.Itoa(controller.Port))
	if err != nil {
		panic("failed to connect to Kafka controller: " + err.Error())
	}

	topicConfig := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(topicConfig...)
	if err != nil {
		if !isTopicAlreadyExists(err) {
			panic("failed to create topic: " + err.Error())
		}
	}
}

func isTopicAlreadyExists(err error) bool {
	return err != nil && (err.Error() == "topic already exists")
}
