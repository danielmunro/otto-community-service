package controller

import (
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	"github.com/google/uuid"
	"net/http"
)

// CreatePostReportV1 - report a post
func CreatePostReportV1(w http.ResponseWriter, r *http.Request) {
	newReport := model.DecodeRequestToNewPostReport(r)
	userUuid := uuid.MustParse(newReport.User.Uuid)
	service.CreateDefaultAuthService().
		DoWithValidSessionAndUser(w, r, userUuid, func() (interface{}, error) {
			return service.CreateDefaultReportService().CreatePostReport(newReport)
		})
}

// CreateReplyReportV1 - report a reply
func CreateReplyReportV1(w http.ResponseWriter, r *http.Request) {
	newReport := model.DecodeRequestToNewPostReport(r)
	userUuid := uuid.MustParse(newReport.User.Uuid)
	service.CreateDefaultAuthService().
		DoWithValidSessionAndUser(w, r, userUuid, func() (interface{}, error) {
			return service.CreateDefaultReportService().CreateReplyReport(newReport)
		})
}
