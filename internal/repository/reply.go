package repository

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ReplyRepository struct {
	conn *gorm.DB
}

func CreateReplyRepository(conn *gorm.DB) *ReplyRepository {
	return &ReplyRepository{conn}
}

func (r *ReplyRepository) Create(reply *entity.Reply) {
	r.conn.Create(reply)
}

func (r *ReplyRepository) FindRepliesForPost(post *entity.Post) []*entity.Reply {
	var replies []*entity.Reply
	r.conn.Where("post_id = ?", post.ID).
		Order("id desc").
		Find(&replies)
	return replies
}

func (r *ReplyRepository) FindOneByUuid(uuid uuid.UUID) (*entity.Reply, error) {
	reply := &entity.Reply{}
	r.conn.Where("uuid = ?", uuid).Find(reply)
	if reply.ID == 0 {
		return nil, errors.New(constants.ErrorMessagePostNotFound)
	}
	return reply, nil
}
