package entity

import (
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Text       string
	UserID     uint
	User       *User
	Visibility model.Visibility `gorm:"default:'public'"`
	Uuid       *uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()"`
	Reports    []*Report        `gorm:"polymorphic:Reported;"`
	Images     []*Image
}

func CreatePost(user *User, post *model.NewPost) *Post {
	if post.Visibility == "" {
		post.Visibility = model.PUBLIC
	}
	return &Post{
		UserID:     user.ID,
		User:       user,
		Text:       post.Text,
		Visibility: post.Visibility,
	}
}
