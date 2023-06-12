package models

import "time"

type Reply struct {
	ReplyID   uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	PostID    uint
	UserID    uint
	Message   string
	Picture   string

	User       User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Post       Post        `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReplyLikes []ReplyLike `gorm:"foreignKey:ReplyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
