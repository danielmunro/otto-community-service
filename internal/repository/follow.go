package repository

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type FollowRepository struct {
	conn *gorm.DB
}

func CreateFollowRepository(conn *gorm.DB) *FollowRepository {
	return &FollowRepository{conn}
}

func (f *FollowRepository) Create(entity *entity.Follow) {
	f.conn.Create(entity)
}

func (f *FollowRepository) FindByUser(user *entity.User) []*entity.Follow {
	var follows []*entity.Follow
	f.conn.Preload("Following").Where("user_id = ?", user.ID).Find(&follows)
	return follows
}

func (f *FollowRepository) FindOne(followUuid uuid.UUID) *entity.Follow {
	var follow *entity.Follow
	f.conn.Preload("Following").Where("uuid = ?", followUuid.String()).Find(&follow)
	return follow
}

func (f *FollowRepository) Update(follow *entity.Follow) {
	f.conn.Save(follow)
}
