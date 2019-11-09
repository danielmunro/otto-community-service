package service_test

import (
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	"github.com/danielmunro/otto-community-service/internal/test"
	"github.com/google/uuid"
	"testing"
)

func Test_PostReport_HappyPath(t *testing.T) {
	// setup
	user1 := createTestUser()
	postService := service.CreateDefaultPostService()
	post, _ := postService.CreatePost(model.CreateNewPost(user1.Uuid, ""))
	user2 := createTestUser()
	reportService := service.CreateDefaultReportService()

	// when
	report, err := reportService.CreatePostReport(model.CreateNewPostReport(user2.Uuid, post.Uuid, "this is offensive"))

	// then
	test.Assert(t, err == nil)
	test.Assert(t, report != nil)
}

func Test_PostReport_Fails_WhenPostMissing(t *testing.T) {
	// setup
	postUuid := uuid.New()
	user := createTestUser()
	reportService := service.CreateDefaultReportService()

	// when
	report, err := reportService.CreatePostReport(model.CreateNewPostReport(user.Uuid, &postUuid, "this is offensive"))

	// then
	test.Assert(t, err != nil)
	test.Assert(t, err.Error() == constants.ErrorMessagePostNotFound)
	test.Assert(t, report == nil)
}

func Test_PostReport_Fails_WhenUserMissing(t *testing.T) {
	// setup
	postUuid := uuid.New()
	userUuid := uuid.New()
	reportService := service.CreateDefaultReportService()

	// when
	report, err := reportService.CreatePostReport(model.CreateNewPostReport(&userUuid, &postUuid, "this is offensive"))

	// then
	test.Assert(t, err != nil)
	test.Assert(t, err.Error() == constants.ErrorMessageUserNotFound)
	test.Assert(t, report == nil)
}

func Test_ReplyReport_HappyPath(t *testing.T) {
	// setup
	user1 := createTestUser()
	postService := service.CreateDefaultPostService()
	post, _ := postService.CreatePost(model.CreateNewPost(user1.Uuid, ""))
	replyService := service.CreateDefaultReplyService()
	reply, _ := replyService.CreateReply(model.CreateNewReply(user1.Uuid, post.Uuid, "test message"))
	user2 := createTestUser()
	reportService := service.CreateDefaultReportService()

	// when
	report, err := reportService.CreateReplyReport(model.CreateNewReplyReport(user2.Uuid, reply.Uuid, "this is offensive"))

	// then
	test.Assert(t, err == nil)
	test.Assert(t, report != nil)
}
