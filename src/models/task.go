package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Completed   bool      `json:"completed" gorm: "default: false"`
	CompletedAt time.Time `json:"completed_at"`
}
