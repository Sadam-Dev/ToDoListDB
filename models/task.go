package models

import "time"

type Task struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	Status      bool
	IsDeleted   bool
	CreatedAt   time.Time
}
