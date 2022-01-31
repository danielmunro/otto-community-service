package service_test

import (
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/service"
	"github.com/danielmunro/otto-community-service/internal/test"
	"github.com/google/uuid"
	"testing"
)

const NumberOfRepliesToCreate = 5

func createReplyModel(post *model.Post, user *entity.User) *model.NewReply {
	return &model.NewReply{
		Post: *post,
		User: model.User{Uuid: user.Uuid.String()},
		Text: "this is a reply",
	}
}

func Test_GetReplies_ForPost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()
	replyService := service.CreateDefaultReplyService()
	post, err := postService.CreatePost(model.CreateNewPost(testUser.Uuid, "this is a test"))

	// expect
	test.Assert(t, err == nil)
	test.Assert(t, post != nil)

	// given
	for i := 0; i < NumberOfRepliesToCreate; i++ {
		_, _ = replyService.CreateReply(createReplyModel(post, testUser))
	}

	// when
	replies, _ := replyService.GetRepliesForPost(uuid.MustParse(post.Uuid))

	// then
	test.Assert(t, len(replies) == NumberOfRepliesToCreate)
}

func Test_CreateReply_Fails_WithMissing_User(t *testing.T) {
	// setup
	testUser := test.CreateTestUser()
	replyService := service.CreateDefaultReplyService()

	// when
	postUuid := uuid.New()
	response, err := replyService.CreateReply(&model.NewReply{
		Post: model.Post{
			Uuid: postUuid.String(),
			Text: "",
		},
		User: model.User{Uuid: testUser.Uuid},
		Text: "this is a reply",
	})

	// then
	test.Assert(t, err != nil)
	test.Assert(t, err.Error() == constants.ErrorMessageUserNotFound)
	test.Assert(t, response == nil)
}

func Test_CreateReply_Fails_WithMissing_Post(t *testing.T) {
	// setup
	testUser := createTestUser()
	replyService := service.CreateDefaultReplyService()

	// when
	postUuid := uuid.New()
	response, err := replyService.CreateReply(&model.NewReply{
		Post: model.Post{
			Uuid: postUuid.String(),
			Text: "",
		},
		User: model.User{Uuid: testUser.Uuid.String()},
		Text: "this is a reply",
	})

	// then
	test.Assert(t, err != nil)
	test.Assert(t, err.Error() == constants.ErrorMessagePostNotFound)
	test.Assert(t, response == nil)
}

func Test_GetReplies_FailsWithMissing_Post(t *testing.T) {
	// setup
	replyService := service.CreateDefaultReplyService()

	// when
	response, err := replyService.GetRepliesForPost(uuid.New())

	// then
	test.Assert(t, err != nil)
	test.Assert(t, err.Error() == constants.ErrorMessagePostNotFound)
	test.Assert(t, response == nil)
}
