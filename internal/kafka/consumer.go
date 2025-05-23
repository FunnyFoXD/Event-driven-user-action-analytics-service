package kafka

import (
	"context"
	"encoding/json"

	"event-tracker/internal/model"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EventPayload struct {
	UserID string `json:"user_id"`
	Action string `json:"action"`
}

func StartConsumer(broker, topic string, db *gorm.DB, logger *zap.Logger) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker},
		Topic:       topic,
		GroupID:     "event-consumer-group",
		StartOffset: kafka.LastOffset,
	})

	go func() {
		defer r.Close()

		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				logger.Error("Error reading message", zap.Error(err))
				continue
			}

			var event model.Event
			if err := json.Unmarshal(m.Value, &event); err != nil {
				logger.Warn("Failed to decode message", zap.ByteString("value", m.Value))
				continue
			}

			if err := db.Create(&event).Error; err != nil {
				logger.Error("Failed to save event", zap.Error(err))
				continue
			}

			logger.Info("Event saved",
				zap.String("user_id", event.UserID),
				zap.String("action", event.Action),
			)
		}
	}()
}
