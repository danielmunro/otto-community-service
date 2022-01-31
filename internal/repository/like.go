package repository

import "github.com/jinzhu/gorm"

type LikeRepository struct {
	conn *gorm.DB
}

func CreateLikeRepository(conn *gorm.DB) *LikeRepository {
	return &LikeRepository{conn}
}
