package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"event-tracker/internal/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func MakeLogsHandler(db *gorm.DB, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")

		var events []model.Event
		query := db.Order("created_at desc")

		if from != "" && to != "" {
			fromTime, err1 := time.Parse(time.RFC3339, from)
			toTime, err2 := time.Parse(time.RFC3339, to)
			if err1 == nil && err2 == nil {
				query = query.Where("created_at BETWEEN ? AND ?", fromTime, toTime)
			}
		}

		if err := query.Find(&events).Error; err != nil {
			logger.Error("Failed to fetch logs", zap.Error(err))
			http.Error(w, "failed to fetch logs", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	}
}
