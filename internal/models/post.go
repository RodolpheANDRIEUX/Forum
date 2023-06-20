package models

import (
	"encoding/base64"
	"time"
)

type Post struct {
	PostID    uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint `gorm:"not null"`
	User      User
	Message   string
	Picture   []byte
	Topic     string
	Like      int `json:"Like"`    // provisoire
	Comment   int `json:"Comment"` // provisoire
}

func (p Post) FormattedCreatedAt() string {
	return p.CreatedAt.Format("02-01-2006 15:04")
}

func (p Post) EncodedImage() string {
	return base64.StdEncoding.EncodeToString(p.Picture)
}
