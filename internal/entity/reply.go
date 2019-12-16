package entity

import (
	"github.com/danielmunro/otto-community-service/internal/enum"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Reply struct {
	gorm.Model
	Text       string
	UserID     uint
	User *User
	Visibility enum.Visibility
	Uuid   *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	PostID uint
	Reports []*Report `gorm:"polymorphic:Reported;"`
}

func CreateReply(user *User, post *Post, reply *model.NewReply) *Reply {
	return &Reply{
		Text:       reply.Text,
		UserID:     user.ID,
		Visibility: enum.PUBLIC,
		PostID: post.ID,
	}
}
