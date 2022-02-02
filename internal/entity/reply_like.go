package entity

type ReplyLike struct {
	ID      uint `gorm:"primary_key"`
	UserID  uint
	User    *User
	ReplyID uint
	Reply   *Reply
}
