package models

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	IsDone      bool
	IsDeleted   bool
	Priority    int
	CreatedAt   time.Time
}
