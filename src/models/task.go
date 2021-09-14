package models

import (
	"time"
)

type Task struct {
	Id          int32     `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Completed   bool      `json:"completed"`
	CompletedAt time.Time `json:"completed_at"`
}
