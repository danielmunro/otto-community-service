package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	iUuid "github.com/danielmunro/otto-community-service/internal/uuid"
	"github.com/google/uuid"
	"net/http"
)

// CreateAReplyV1 - create a reply
func CreateAReplyV1(w http.ResponseWriter, r *http.Request) {
	newReplyModel := model.DecodeRequestToNewReply(r)
	userUuid := uuid.MustParse(newReplyModel.Message.User.Uuid)
	service.CreateDefaultAuthService().
		DoWithValidSessionAndUser(w, r, userUuid, func() (interface{}, error) {
			return service.CreateDefaultReplyService().CreateReply(newReplyModel)
		})
}

func GetPostRepliesV1(w http.ResponseWriter, r *http.Request) {
	postUuid := iUuid.GetUuidFromPathSecondPosition(r.URL.Path)
	replies, err := service.CreateDefaultReplyService().GetRepliesForPost(postUuid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(replies)
	_, _ = w.Write(data)
}
