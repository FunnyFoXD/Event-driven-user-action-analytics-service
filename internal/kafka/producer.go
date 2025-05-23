package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
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
