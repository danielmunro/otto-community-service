package repository

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type PostRepository struct {
	conn *gorm.DB
}

func CreatePostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{conn}
}

func (p *PostRepository) Save(entity *entity.Post) {
	p.conn.Save(entity)
}

func (p *PostRepository) FindByUser(user *entity.User, limit int) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Where("user_id = ? and deleted_at IS NULL", user.ID).
		Order("id desc").
		Limit(limit).
		Find(&posts)
	return posts
}

func (p *PostRepository) FindOneByUuid(uuid uuid.UUID) (*entity.Post, error) {
	post := &entity.Post{}
	p.conn.
		Preload("User").
		Where("uuid = ? and deleted_at IS NULL", uuid).
		Find(post)
	if post.ID == 0 {
		return nil, errors.New(constants.ErrorMessagePostNotFound)
	}
	return post, nil
}

func (p *PostRepository) FindByUserFollows(userUuid uuid.UUID, limit int) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Table("posts").
		Joins("join follows on follows.user_id = posts.user_id").
		Joins("join users on follows.following_id = users.id").
		Where("users.uuid = ? and posts.deleted_at IS NULL", userUuid.String()).
		Order("id desc").
		Limit(limit).
		Find(&posts)
	return posts
}

func (p *PostRepository) FindAll(limit int) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Table("posts").
		Joins("join users on posts.user_id = users.id").
		Where("posts.deleted_at is null and users.deleted_at is null").
		Order("id desc").
		Limit(limit).
		Find(&posts)
	return posts
}
