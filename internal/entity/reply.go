package entity

import (
	"github.com/danielmunro/otto-community-service/internal/enum"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Reply struct {
	gorm.Model
	Message
	Uuid   *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	PostID uint
	Reports []*Report `gorm:"polymorphic:Reported;"`
}

func CreateReply(user *User, post *Post, reply *model.NewReply) *Reply {
	return &Reply{
		Message: Message{
			Text:       reply.Message.Text,
			UserID:     user.ID,
			Visibility: enum.PUBLIC,
		},
		PostID: post.ID,
	}
}
