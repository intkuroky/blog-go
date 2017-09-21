package models

import "time"

type Article struct {
	ID        uint64 `gorm:"column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Article) TableName() string {
	return "articles"
}
