package models

import "time"

type Post struct {
	PostID    uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	Message   string
	Picture   []byte
	Topic     string

type PostWeb struct {
	PostID        uint   `gorm:"primaryKey;autoIncrement" json:"post_id"`
	UserID        uint   `json:"user_id"`
	Username      string `json:"username"`
	ProfilPicture string `json:"profilPicture"`
	Message       string `json:"message"`
	Picture       string `json:"picture"`
	Topic         string `json:"topic"`
	Reply         int    `json:"reply"`
	Like          int    `json:"like"`

	//User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//Replies []Reply `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
