package models

import "time"

type PostLike struct {
	LikeID    uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	PostID    uint
	Reaction  int
	User      User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Post      Post `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ReplyLike struct {
	LikeID    uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	ReplyID   uint
	Reaction  int

	User  User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Reply Reply `gorm:"foreignKey:ReplyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
