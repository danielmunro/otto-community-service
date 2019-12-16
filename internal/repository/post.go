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

func (p *PostRepository) FindByUser(user *entity.User) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Where("user_id = ?", user.ID).
		Order("id desc").
		Limit(constants.UserPostsDefaultPageSize).
		Find(&posts)
	return posts
}

func (p *PostRepository) FindOneByUuid(uuid uuid.UUID) (*entity.Post, error) {
	post := &entity.Post{}
	p.conn.
		Preload("User").
		Where("uuid = ?", uuid).
		Find(post)
	if post.ID == 0 {
		return nil, errors.New(constants.ErrorMessagePostNotFound)
	}
	return post, nil
}

func (p *PostRepository) FindByUserFollows(userUuid uuid.UUID) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Table("posts").
		Joins("join follows on follows.user_id = posts.user_id").
		Joins("join users on follows.following_id = users.id").
		Where("users.uuid = ?", userUuid.String()).
		Order("id desc").
		Find(&posts)
	return posts
}
