package models

type User struct {
	Model
	Name		string	`gorm:"not null; size:64"`
	Password	string	`gorm:"not null; size:128"`
	GroupId		int8	`gorm:"default 0"`
	Score		int32	`gorm:"default 0"`
}

