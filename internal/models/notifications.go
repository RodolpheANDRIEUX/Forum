package models

import (
	"time"
)

type Notifications struct {
	NotificationID uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         uint `gorm:"not null"`
	Message        string
	User           User
	Deleted        bool `gorm:"default:false"`
}

func (p Notifications) FormatDateNotif() string {
	return p.CreatedAt.Format("02-01-2006 15:04")
}
