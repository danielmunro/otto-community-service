package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	iUuid "github.com/danielmunro/otto-community-service/internal/uuid"
	"github.com/google/uuid"
	"net/http"
)

// CreateNewPostV1 - create a new post
func CreateNewPostV1(w http.ResponseWriter, r *http.Request) {
	newPostModel := model.DecodeRequestToNewPost(r)
	userUuid := uuid.MustParse(newPostModel.Message.User.Uuid)
	service.CreateDefaultAuthService().
		DoWithValidSessionAndUser(w, r, userUuid, func() (interface{}, error) {
			return service.CreateDefaultPostService().CreatePost(newPostModel)
		})
}

// GetPostV1 - get a post
func GetPostV1(w http.ResponseWriter, r *http.Request) {
	post, err := service.CreateDefaultPostService().GetPost(iUuid.GetUuidFromPath(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(post)
	_, _ = w.Write(data)
}
