package entity

import "github.com/jinzhu/gorm"

type PostLike struct {
	gorm.Model
	UserID uint
	User   *User
	PostID uint
	Post   *Post
}
