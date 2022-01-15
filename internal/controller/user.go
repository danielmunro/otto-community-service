package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/service"
	"github.com/danielmunro/otto-community-service/internal/uuid"
	"net/http"
)

// GetUserPostsV1 - get posts by a user
func GetUserPostsV1(w http.ResponseWriter, r *http.Request) {
	posts, _ := service.CreateDefaultPostService().GetPostsForUser(uuid.GetUuidFromPathSecondPosition(r.URL.Path))
	data, _ := json.Marshal(posts)
	_, _ = w.Write(data)
}

// GetUserV1 - get a user
func GetUserV1(w http.ResponseWriter, r *http.Request) {
	user, err := service.CreateDefaultUserService().GetUser(uuid.GetUuidFromPathSecondPosition(r.URL.Path))
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
		GetSuggestedFollowsForUser(uuid.GetUuidFromPathSecondPosition(r.URL.Path))
	data, _ := json.Marshal(users)
	_, _ = w.Write(data)
}
