package models

import "time"

type Reply struct {
	ReplyID    uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PostID     uint
	UserID     uint
	User       User
	Message    string
	Picture    []byte
	ReplyLikes []ReplyLike `gorm:"foreignKey:ReplyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
