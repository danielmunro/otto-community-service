package mapper

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
)

func GetReplyModelFromEntity(reply *entity.Reply) *model.Reply {
	return &model.Reply{
		Uuid:      reply.Uuid.String(),
		CreatedAt: reply.CreatedAt,
		Text:      reply.Text,
		User:      *GetUserModelFromEntity(reply.User),
		Likes:     reply.Likes,
	}
}

func GetReplyModelsFromEntities(replies []*entity.Reply) []*model.Reply {
	replyModels := make([]*model.Reply, len(replies))
	for i, v := range replies {
		replyModels[i] = GetReplyModelFromEntity(v)
	}
	return replyModels
}
