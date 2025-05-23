package model

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UserID string `json:"user_id"`
	Action string `json:"action"`
}
