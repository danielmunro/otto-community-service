package entity

import (
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Message
	Uuid *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Reports []*Report `gorm:"polymorphic:Reported;"`
}

func CreatePost(user *User, post *model.NewPost) *Post {
	return &Post{
		Message: Message{
			UserID: user.ID,
			Text:   post.Message.Text,
		},
	}
}
