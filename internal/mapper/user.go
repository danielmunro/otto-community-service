package mapper

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"time"
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
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email,
		ProfilePic: user.ProfilePic,
		Location:   user.Location,
		BioMessage: user.BioMessage,
		Birthday:   user.Birthday.String(),
		CreatedAt:  user.CreatedAt,
	}
}

func GetUserEntityFromModel(user *model.User) *entity.User {
	userUuid := uuid.MustParse(user.Uuid)
	birthday, _ := time.Parse("2006-01-02", user.Birthday)
	return &entity.User{
		Uuid: &userUuid,
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
		ProfilePic: user.ProfilePic,
		Location: user.Location,
		BioMessage: user.BioMessage,
		Birthday: birthday,
	}
}
