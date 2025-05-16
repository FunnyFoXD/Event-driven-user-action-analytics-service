package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Event struct {
	UserId string `json:"user_id"`
	Action string `json:"action"`
}

func MakeEventHandler(logger *zap.Logger) http.HandlerFunc {
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

		logger.Info("Event recieved",
			zap.String("user_id", evt.UserId),
			zap.String("action", evt.Action))

		w.WriteHeader(http.StatusAccepted)
	}
}
