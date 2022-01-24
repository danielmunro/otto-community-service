package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/auth/model"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/service"
	iUuid "github.com/danielmunro/otto-community-service/internal/uuid"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CreateNewFollowV1 - create a new follow
func CreateNewFollowV1(w http.ResponseWriter, r *http.Request) {
	newFollowModel := mapper.DecodeRequestToNewFollow(r)
	userUuid := iUuid.GetUuidFromPathSecondPosition(r.URL.Path)
	service.CreateDefaultAuthService().
		DoWithValidSessionAndUser(w, r, userUuid, func() (interface{}, error) {
			return service.CreateDefaultFollowService().CreateFollow(userUuid, newFollowModel)
		})
}

// GetUserFollowersByUsernameV1 - get user followers
func GetUserFollowersByUsernameV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameParam := params["username"]

	follows, err := service.CreateDefaultFollowService().GetUserFollowersByUsername(usernameParam)
	if err != nil {
		log.Print("error received from get user follows :: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(follows)
	_, _ = w.Write(data)
}

// GetUserFollowersV1 - get user followers
func GetUserFollowersV1(w http.ResponseWriter, r *http.Request) {
	follows, err := service.CreateDefaultFollowService().GetUserFollowers(iUuid.GetUuidFromPathSecondPosition(r.URL.Path))
	if err != nil {
		log.Print("error received from get user follows :: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(follows)
	_, _ = w.Write(data)
}

// GetUserFollowsByUsernameV1 - get user follows
func GetUserFollowsByUsernameV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameParam := params["username"]

	follows, err := service.CreateDefaultFollowService().GetUserFollowsByUsername(usernameParam)
	if err != nil {
		log.Print("error received from get user follows :: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(follows)
	_, _ = w.Write(data)
}

// GetUserFollowsV1 - get user follows
func GetUserFollowsV1(w http.ResponseWriter, r *http.Request) {
	follows, err := service.CreateDefaultFollowService().GetUserFollows(iUuid.GetUuidFromPathSecondPosition(r.URL.Path))
	if err != nil {
		log.Print("error received from get user follows :: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(follows)
	_, _ = w.Write(data)
}

// DeleteFollowV1 - delete a follow
func DeleteFollowV1(w http.ResponseWriter, r *http.Request) {
	followUuid := iUuid.GetUuidFromPathSecondPosition(r.URL.Path)
	service.CreateDefaultAuthService().DoWithValidSession(w, r, func(session *model.Session) (interface{}, error) {
		err := service.CreateDefaultFollowService().DeleteFollow(followUuid, uuid.MustParse(session.User.Uuid))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return nil, err
	})
}
