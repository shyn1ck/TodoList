package models

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(100)"`
	Description string    `json:"description" gorm:"type:text"`
	IsDone      bool      `json:"is_done" gorm:"column:is_done;type:boolean;default:false"`
	Priority    int       `json:"priority" gorm:"type:integer;default:0"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:timestamp with time zone;default:now()"`
	IsDeleted   bool      `json:"is_delete" gorm:"column:is_deleted;type:boolean;default:false"`
}
