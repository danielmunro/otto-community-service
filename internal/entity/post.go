package entity

import (
	"github.com/danielmunro/otto-community-service/internal/enum"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Text       string
	UserID     uint
	User *User
	Visibility enum.Visibility
	Uuid *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Reports []*Report `gorm:"polymorphic:Reported;"`
}

func CreatePost(user *User, post *model.NewPost) *Post {
	return &Post{
		UserID: user.ID,
		User: user,
		Text:   post.Text,
	}
}
