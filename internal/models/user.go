package models

import (
	"time"
)

type User struct {
	UserID        uint `gorm:"primaryKey;unique"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Role          string
	Username      string `gorm:"unique"`
	Email         string `gorm:"unique"`
	ProfileImg    []byte
	Password      string
	Notifications []Notifications `gorm:"foreignKey:UserID"`
}

func (p User) FormattedDate() string {
	return p.CreatedAt.Format("02-01-2006 15:04")
}
