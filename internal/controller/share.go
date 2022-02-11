package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// GetShareV1 - get a share
func GetShareV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuidParam := uuid.MustParse(params["uuid"])
	share, err := service.CreateDefaultShareService().GetShare(uuidParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(share)
	_, _ = w.Write(data)
}

// CreateShareV1 - create a reshare
func CreateShareV1(w http.ResponseWriter, r *http.Request) {
	newShareParam := model.DecodeRequestToNewShare(r)
	share, err := service.CreateDefaultShareService().CreateShare(newShareParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(share)
	_, _ = w.Write(data)
}
