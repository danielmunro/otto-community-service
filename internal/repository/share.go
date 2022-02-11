package repository

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ShareRepository struct {
	conn *gorm.DB
}

func CreateShareRepository(conn *gorm.DB) *ShareRepository {
	return &ShareRepository{conn}
}

func (s *ShareRepository) Save(entity *entity.Post) {
	s.conn.Save(entity)
}

func (s *ShareRepository) FindOneByUuid(shareUuid uuid.UUID) (*entity.Post, error) {
	sharePost := &entity.Post{}
	s.conn.Preload("SharePost").
		Preload("User").
		Where("uuid = ?", shareUuid).
		Find(sharePost)
	if sharePost.ID == 0 {
		return nil, errors.New("no share found")
	}
	return sharePost, nil
}
