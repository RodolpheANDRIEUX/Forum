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
	Report    uint `gorm:"default:0"`
}

func (p Post) FormattedCreatedAt() string {
	return p.CreatedAt.Format("02-01-2006 15:04")
}

func (p Post) EncodedImage() string {
	return base64.StdEncoding.EncodeToString(p.Picture)
}

type PostWeb struct {
	PostID         uint   ` json:"post_id"`
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	ProfilePicture []byte `json:"profilPicture"`
	Message        string `json:"message"`
	Picture        []byte `json:"picture"`
	Topic          string `json:"topic"`
	Reply          int    `json:"reply"`
	Like           int    `json:"like"`
}
