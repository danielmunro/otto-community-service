package entity

import (
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Text          string
	UserID        uint
	User          *User
	Visibility    model.Visibility `gorm:"default:'public'"`
	Uuid          *uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()"`
	Reports       []*Report        `gorm:"polymorphic:Reported;"`
	Images        []*Image
	Likes         uint
	Replies       uint
	ThreadPostID  uint `gorm:"foreignkey:Post"`
	ThreadPost    *Post
	ReplyToPostID uint `gorm:"foreignkey:Post"`
	ReplyToPost   *Post
	SharePostID   uint `gorm:"foreignkey:Post"`
	SharePost     *Post
}

func CreatePost(user *User, post *model.NewPost) *Post {
	if post.Visibility == "" {
		post.Visibility = model.PUBLIC
	}
	postUuid := uuid.New()
	return &Post{
		Uuid:       &postUuid,
		UserID:     user.ID,
		User:       user,
		Text:       post.Text,
		Visibility: post.Visibility,
	}
}

func CreateReply(user *User, post *Post, reply *model.NewReply) *Post {
	return &Post{
		Text:          reply.Text,
		UserID:        user.ID,
		Visibility:    model.PUBLIC,
		ReplyToPost:   post,
		ReplyToPostID: post.ID,
		User:          user,
	}
}
