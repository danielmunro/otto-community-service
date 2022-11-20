package service_test

import (
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	"github.com/danielmunro/otto-community-service/internal/test"
	"testing"
)

func Test_ShareService_CanCreate_NewShare(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreatePostService()
	shareService := service.CreateDefaultShareService()
	post, _ := postService.CreatePost(*testUser.Uuid, model.CreateNewPost(testUser.Uuid, message))
	newShare := &model.NewShare{
		Text: "Yo",
		User: model.User{
			Uuid: testUser.Uuid.String(),
		},
		Post: *post,
	}

	// when
	share, err := shareService.CreateShare(newShare)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, share != nil)
}
