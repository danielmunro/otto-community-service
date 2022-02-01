package entity

import "github.com/jinzhu/gorm"

type PostLike struct {
	gorm.Model
	UserID uint `gorm:"unique_index:unique_user_post"`
	User   *User
	PostID uint `gorm:"unique_index:unique_user_post"`
	Post   *Post
}
