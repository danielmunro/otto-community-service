package repository

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"strings"
)

type LikeRepository struct {
	conn *gorm.DB
}

func CreateLikeRepository(conn *gorm.DB) *LikeRepository {
	return &LikeRepository{conn}
}

func (l *LikeRepository) FindLikesForPosts(postUuids []uuid.UUID) []*entity.PostLike {
	query := "SELECT * " +
		"FROM post_likes " +
		"WHERE post_id IN (SELECT id FROM posts WHERE uuid IN ("
	for range postUuids {
		query += "?, "
	}
	query = strings.TrimRight(query, ", ")
	query += "))"
	var postLikes []*entity.PostLike
	l.conn.Raw(query, postUuids).Preload("Posts").Scan(&postLikes)
	return postLikes
}
