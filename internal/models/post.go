package models

import (
	"encoding/base64"
	"time"
)

type Post struct {
	PostID    uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool `gorm:"default:false"`
	UserID    uint `gorm:"not null"`
	User      User
	Message   string
	Picture   []byte
	Topic     string
	Replies   []Reply
	Like      int  `json:"Like"`    // provisoire
	Comment   int  `json:"Comment"` // provisoire
	Report    uint `gorm:"default:0"`
}

func (p Post) FormattedCreatedAt() string {
	return p.CreatedAt.Format("02-01-2006 15:04")
}

func (p Post) EncodedImage() string {
	return base64.StdEncoding.EncodeToString(p.Picture)
}
