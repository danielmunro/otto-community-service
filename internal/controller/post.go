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
	userUuid := uuid.MustParse(newPostModel.User.Uuid)
	service.CreateDefaultAuthService().
		DoWithValidSessionAndUser(w, r, userUuid, func() (interface{}, error) {
			return service.CreateDefaultPostService().CreatePost(newPostModel)
		})
}

// GetPostV1 - get a post
func GetPostV1(w http.ResponseWriter, r *http.Request) {
	post, err := service.CreateDefaultPostService().GetPost(iUuid.GetUuidFromPathSecondPosition(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(post)
	_, _ = w.Write(data)
}

// GetUserFollowsPostsV1 - get a user's friend's posts
func GetUserFollowsPostsV1(w http.ResponseWriter, r *http.Request) {
	posts, err := service.CreateDefaultPostService().GetPostsForUserFollows(
		iUuid.GetUuidFromPathThirdPosition(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(posts)
	_, _ = w.Write(data)
}

func GetNewPostsV1(w http.ResponseWriter, r *http.Request) {
	posts := service.CreateDefaultPostService().GetNewPosts(iUuid.GetUuidFromPathSecondPosition(r.URL.Path))
	data, _ := json.Marshal(posts)
	_, _ = w.Write(data)
}

// GetPosts - get posts
func GetPostsV1(w http.ResponseWriter, r *http.Request) {
	authService := service.CreateDefaultAuthService()
	sessionToken := authService.GetSessionToken(r)
	if sessionToken == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	session, _ := authService.GetSession(sessionToken)
	posts, _ := service.CreateDefaultPostService().GetPosts(uuid.MustParse(session.User.Uuid))
	data, _ := json.Marshal(posts)
	_, _ = w.Write(data)
}
