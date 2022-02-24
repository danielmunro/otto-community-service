package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/service"
	uuid2 "github.com/danielmunro/otto-community-service/internal/uuid"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// GetUserPostsV1 - get posts by a user
func GetUserPostsV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=30")
	params := mux.Vars(r)
	username := params["username"]
	session := service.CreateDefaultAuthService().GetSessionFromRequest(r)
	var viewerUuid uuid.UUID
	if session != nil {
		viewerUuid = uuid.MustParse(session.User.Uuid)
	}
	posts, _ := service.CreateDefaultPostService().GetPostsForUser(
		username, &viewerUuid, constants.UserPostsDefaultPageSize)
	data, _ := json.Marshal(posts)
	_, _ = w.Write(data)
}

// GetUserV1 - get a user
func GetUserV1(w http.ResponseWriter, r *http.Request) {
	user, err := service.CreateDefaultUserService().GetUser(uuid2.GetUuidFromPathSecondPosition(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(user)
	_, _ = w.Write(data)
}

// GetSuggestedFollowsForUserV1 - Get suggested follows for user
func GetSuggestedFollowsForUserV1(w http.ResponseWriter, r *http.Request) {
	users := service.CreateDefaultUserService().
		GetSuggestedFollowsForUser(uuid2.GetUuidFromPathSecondPosition(r.URL.Path))
	data, _ := json.Marshal(users)
	_, _ = w.Write(data)
}
