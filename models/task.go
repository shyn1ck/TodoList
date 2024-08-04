package models

import (
	"time"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
	IsDone      bool
	Priority    int
	CreatedAt   time.Time
}
