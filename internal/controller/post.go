package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	iUuid "github.com/danielmunro/otto-community-service/internal/uuid"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
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
	authService := service.CreateDefaultAuthService()
	session := authService.GetSessionFromRequest(r)
	var viewerUuid uuid.UUID
	if session != nil {
		viewerUuid = uuid.MustParse(session.User.Uuid)
	}
	post, err := service.CreateDefaultPostService().GetPost(
		&viewerUuid,
		iUuid.GetUuidFromPathSecondPosition(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(post)
	_, _ = w.Write(data)
}

// GetUserFollowsPostsV1 - get a user's friend's posts
func GetUserFollowsPostsV1(w http.ResponseWriter, r *http.Request) {
	limit := constants.UserPostsDefaultPageSize
	posts, err := service.CreateDefaultPostService().GetPostsForUserFollows(
		iUuid.GetUuidFromPathThirdPosition(r.URL.Path), limit)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(posts)
	_, _ = w.Write(data)
}

func GetNewPostsV1(w http.ResponseWriter, r *http.Request) {
	limit := constants.UserPostsDefaultPageSize
	posts := service.CreateDefaultPostService().GetNewPosts(
		iUuid.GetUuidFromPathSecondPosition(r.URL.Path), limit)
	data, _ := json.Marshal(posts)
	_, _ = w.Write(data)
}

// GetPostsV1 - get posts
func GetPostsV1(w http.ResponseWriter, r *http.Request) {
	authService := service.CreateDefaultAuthService()
	session := authService.GetSessionFromRequest(r)
	var viewerUuid uuid.UUID
	if session != nil {
		viewerUuid = uuid.MustParse(session.User.Uuid)
	}
	limit := constants.UserPostsDefaultPageSize
	posts, _ := service.CreateDefaultPostService().GetPosts(&viewerUuid, limit)
	data, _ := json.Marshal(posts)
	_, _ = w.Write(data)
}

// DeletePostV1 - delete a post
func DeletePostV1(w http.ResponseWriter, r *http.Request) {
	authService := service.CreateDefaultAuthService()
	session := authService.GetSessionFromRequest(r)
	if session == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	params := mux.Vars(r)
	postUuid, err := uuid.Parse(params["uuid"])
	if err != nil {
		log.Print("malformed uuid for delete post :: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userUuid := uuid.MustParse(session.User.Uuid)
	err = service.CreateDefaultPostService().DeletePost(postUuid, userUuid)
	if err != nil {
		log.Print("delete error, error :: ", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
