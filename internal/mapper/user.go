package mapper

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
)

func GetUserModelFromEntity(user *entity.User) *model.User {
	return &model.User{
		Uuid:      user.Uuid.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}

func GetUserEntityFromModel(user *model.User) *entity.User {
	userUuid := uuid.MustParse(user.Uuid)
	return &entity.User{
		Uuid: &userUuid,
	}
}
