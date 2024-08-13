package models

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Status      bool
	IsDeleted   bool
	CreatedAt   time.Time
}
