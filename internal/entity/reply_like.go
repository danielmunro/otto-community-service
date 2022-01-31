package entity

import "github.com/jinzhu/gorm"

type ReplyLike struct {
	gorm.Model
	UserID  uint
	User    *User
	ReplyID uint
	Reply   *Reply
}
