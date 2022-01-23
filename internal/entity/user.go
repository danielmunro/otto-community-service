package entity

import (
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Uuid       *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username   string     `gorm:"unique;not null"`
	ProfilePic string
	Name       string
	BioMessage string
	Birthday   time.Time
	Follows    []*Follow
	Posts      []*Post
	Replies    []*Reply
}

func (u *User) UpdateUserProfileFromModel(user *model.User) {
	u.Name = user.Name
	u.ProfilePic = user.ProfilePic
	u.BioMessage = user.BioMessage
	u.Birthday = user.Birthday
}
