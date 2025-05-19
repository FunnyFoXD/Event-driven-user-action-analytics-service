package handler

import (
	"encoding/json"
	"event-tracker/internal/kafka"
	"net/http"

	"go.uber.org/zap"
)

type Event struct {
	UserId string `json:"user_id"`
	Action string `json:"action"`
}

func MakeEventHandler(logger *zap.Logger, producer *kafka.Producer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var evt Event
		if err := json.NewDecoder(r.Body).Decode(&evt); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			logger.Warn("Invalid event JSON", zap.Error(err))
			return
		}

		data, err := json.Marshal(evt)
		if err != nil {
			http.Error(w, "failed to serialize event", http.StatusInternalServerError)
			logger.Error("Kafka send error", zap.Error(err))
			return
		}

		if err := producer.SendMessage(r.Context(), evt.UserId, data); err != nil {
			http.Error(w, "failed to send to kafka", http.StatusInternalServerError)
			logger.Error("Kafka send error", zap.Error(err))
			return
		}

		logger.Info("Event recieved",
			zap.String("user_id", evt.UserId),
			zap.String("action", evt.Action))

		w.WriteHeader(http.StatusAccepted)
	}
}
