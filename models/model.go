package models

import "time"

type Model struct {
	ID        uint64 `gorm:"column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}