package mapper

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
)

func GetUserModelsFromEntities(users []*entity.User) []*model.User {
	userModels := make([]*model.User, len(users))
	for i, v := range users {
		userModels[i] = GetUserModelFromEntity(v)
	}
	return userModels
}


func GetUserModelFromEntity(user *entity.User) *model.User {
	return &model.User{
		Uuid:       user.Uuid.String(),
		CreatedAt:  user.CreatedAt,
	}
}

func GetUserEntityFromModel(user *model.User) *entity.User {
	userUuid := uuid.MustParse(user.Uuid)
	return &entity.User{
		Uuid: &userUuid,
	}
}
