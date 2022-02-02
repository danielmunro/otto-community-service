package repository

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/jinzhu/gorm"
)

type LikeRepository struct {
	conn *gorm.DB
}

func CreateLikeRepository(conn *gorm.DB) *LikeRepository {
	return &LikeRepository{conn}
}

func (l *LikeRepository) FindLikesForPosts(postIds []uint) []*entity.PostLike {
	query := "SELECT * " +
		"FROM post_likes " +
		"WHERE post_id IN (?) AND deleted_at IS NULL"
	var postLikes []*entity.PostLike
	l.conn.Raw(query, postIds).Scan(&postLikes)
	return postLikes
}

func (l *LikeRepository) Create(postLike *entity.PostLike) {
	l.conn.Create(postLike)
}

func (l *LikeRepository) Save(postLike *entity.PostLike) {
	l.conn.Save(postLike)
}

func (l *LikeRepository) FindByPostAndUser(post *entity.Post, user *entity.User) (*entity.PostLike, error) {
	postLike := &entity.PostLike{}
	l.conn.Where("user_id = ? and post_id = ? and deleted_at is null", user.ID, post.ID).Find(postLike)
	if postLike.ID == 0 {
		return nil, errors.New("no post like found")
	}
	return postLike, nil
}

func (l *LikeRepository) DeletePostLike(postLike *entity.PostLike) {
	l.conn.Delete(postLike)
}
