package mapper

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
)

func GetPostModelFromEntity(user *entity.User, post *entity.Post) *model.Post {
	return &model.Post{
		Uuid: post.Uuid,
		Message: model.Message{
			Text: post.Text,
			User: model.User{
				Uuid: user.Uuid.String(),
			},
		},
		CreatedAt: post.CreatedAt,
	}
}

func GetPostModelsFromEntities(user *entity.User, posts []*entity.Post) []*model.Post {
	postModels := make([]*model.Post, len(posts))
	for i, v := range posts {
		postModels[i] = GetPostModelFromEntity(user, v)
	}
	return postModels
}
