package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/service"
	"github.com/danielmunro/otto-community-service/internal/uuid"
	"net/http"
)

// CreateNewFollowV1 - create a new follow
func CreateNewFollowV1(w http.ResponseWriter, r *http.Request) {
	newFollowModel := mapper.DecodeRequestToNewFollow(r)
	userUuid := uuid.GetUuidFromPathSecondPosition(r.URL.Path)
	service.CreateDefaultAuthService().
		DoWithValidSessionAndUser(w, r, userUuid, func() (interface{}, error) {
			return service.CreateDefaultFollowService().CreateFollow(userUuid, newFollowModel)
		})
}

// GetUserFollowsV1 - get user follows
func GetUserFollowsV1(w http.ResponseWriter, r *http.Request) {
	follows, err := service.CreateDefaultFollowService().GetUserFollowers(uuid.GetUuidFromPathSecondPosition(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(follows)
	_, _ = w.Write(data)
}
