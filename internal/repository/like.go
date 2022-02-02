package repository

import (
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
		"WHERE post_id IN (?)"
	var postLikes []*entity.PostLike
	l.conn.Raw(query, postIds).Scan(&postLikes)
	return postLikes
}

func (l *LikeRepository) CreatePostLike(postLike *entity.PostLike) {
	l.conn.Create(postLike)
}
