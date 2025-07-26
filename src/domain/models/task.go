package models

import "time"

// Task represents a task in the system.
type Task struct {
	ID          uint64
	Title       string
	Description string
	CreatedAt   time.Time
}
