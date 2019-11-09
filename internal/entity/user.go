package entity

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uuid    *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name    string
	Follows []*Follow
	Posts   []*Post
	Replies []*Reply
}
