package entity

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Uuid    *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name    string
	Username string
	ProfilePic string
	BioMessage string
	Location string
	Email string
	EmailVerified bool
	Phone string
	PhoneVerified bool
	Birthday time.Time
	Follows []*Follow
	Posts   []*Post
	Replies []*Reply
}
