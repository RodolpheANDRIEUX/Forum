package models

import (
	"time"
)

type User struct {
	UserID     uint `gorm:"primaryKey;unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Role       string
	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	ProfileImg []byte
	Password   string

	//Posts []Post `gorm:"foreignKey:UserID"`
	//Like  []Like  `gorm:"foreignKey:UserID"`
	//Reply []Reply `gorm:"foreignKey:UserID"`
}
