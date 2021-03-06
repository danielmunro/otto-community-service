package mapper

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
	"net/http"
)

func DecodeRequestToNewFollow(r *http.Request) *model.NewFollow {
	decoder := json.NewDecoder(r.Body)
	var data *model.NewFollow
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func GetFollowModelFromEntity(follow *entity.Follow, user *entity.User, following *entity.User) *model.Follow {
	return &model.Follow{
		Uuid: follow.Uuid.String(),
		User: model.User{
			Uuid: user.Uuid.String(),
		},
		Following: model.User{
			Uuid: following.Uuid.String(),
		},
	}
}

func GetFollowsModelFromEntities(follows []*entity.Follow, user *entity.User) []*model.Follow {
	followModels := make([]*model.Follow, len(follows))
	for i, v := range follows {
		followModels[i] = &model.Follow{
			Uuid: v.Uuid.String(),
			User: model.User{
				Uuid: user.Uuid.String(),
			},
			Following: model.User{
				Uuid: v.Following.Uuid.String(),
			},
		}
	}
	return followModels
}
