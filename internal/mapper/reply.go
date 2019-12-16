package mapper

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
)

func GetReplyModelFromEntity(reply *entity.Reply) *model.Reply {
	return &model.Reply{
		Uuid: reply.Uuid,
		CreatedAt: reply.CreatedAt,
		Text:       reply.Text,
	}
}

func GetReplyModelsFromEntities(replies []*entity.Reply) []*model.Reply {
	replyModels := make([]*model.Reply, len(replies))
	for i, v := range replies {
		replyModels[i] = GetReplyModelFromEntity(v)
	}
	return replyModels
}
