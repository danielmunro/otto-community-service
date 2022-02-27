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
		Preload("Images").
		Preload("SharePost").
		Where("user_id = ? AND deleted_at IS NULL AND reply_to_post_id = 0", user.ID).
		Order("id desc").
		Limit(limit).
		Find(&posts)
	return posts
}

func (p *PostRepository) FindByLikes(user *entity.User, limit int) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Preload("Images").
		Preload("SharePost").
		Table("posts").
		Joins("join post_likes on post_likes.post_id = posts.id").
		Where("post_likes.user_id = ? AND deleted_at IS NULL AND reply_to_post_id = 0", user.ID).
		Order("id desc").
		Limit(limit).
		Find(&posts)
	return posts
}

func (p *PostRepository) FindOneByUuid(uuid uuid.UUID) (*entity.Post, error) {
	post := &entity.Post{}
	p.conn.
		Preload("User").
		Preload("Images").
		Preload("SharePost").
		Where("uuid = ? AND deleted_at IS NULL", uuid).
		Find(post)
	if post.ID == 0 {
		return nil, errors.New(constants.ErrorMessagePostNotFound)
	}
	return post, nil
}

func (p *PostRepository) FindByUserFollows(username string, limit int) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Preload("Images").
		Preload("SharePost").
		Table("posts").
		Joins("join follows on follows.following_id = posts.user_id").
		Joins("join users on follows.user_id = users.id").
		Where("users.username = ? AND posts.deleted_at IS NULL AND posts.reply_to_post_id = 0", username).
		Order("id desc").
		Limit(limit).
		Find(&posts)
	return posts
}

func (p *PostRepository) FindAll(limit int) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Preload("Images").
		Preload("SharePost").
		Table("posts").
		Joins("join users on posts.user_id = users.id").
		Where("posts.deleted_at IS NULL AND users.deleted_at IS NULL AND posts.reply_to_post_id = 0").
		Order("id desc").
		Limit(limit).
		Find(&posts)
	return posts
}

func (p *PostRepository) FindByIDs(ids []uint) []*entity.Post {
	var posts []*entity.Post
	p.conn.
		Preload("User").
		Preload("Images").
		Preload("SharePost").
		Table("posts").
		Joins("join users on posts.user_id = users.id").
		Where("posts.deleted_at IS NULL AND posts.id IN (?)", ids).
		Order("id desc").
		Find(&posts)
	return posts
}
