package mapper

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
)

func GetPostModelFromEntity(post *entity.Post) *model.Post {
	return &model.Post{
		Uuid: post.Uuid,
		Text: post.Text,
		User: *GetUserModelFromEntity(post.User),
		CreatedAt: post.CreatedAt,
	}
}

func GetPostModelsFromEntities(posts []*entity.Post) []*model.Post {
	postModels := make([]*model.Post, len(posts))
	for i, v := range posts {
		postModels[i] = GetPostModelFromEntity(v)
	}
	return postModels
}
