package models

import "time"

type User struct {
	UserID    uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      string
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string
}
