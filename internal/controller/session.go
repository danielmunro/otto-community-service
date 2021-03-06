package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/auth/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	"net/http"
)

// GetUserV1 - get a user
func CreateSessionV1(w http.ResponseWriter, r *http.Request) {
	newSession := model.DecodeRequestToNewSession(r)
	user, err := service.CreateDefaultSessionService().CreateSession(newSession)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(user)
	_, _ = w.Write(data)
}
