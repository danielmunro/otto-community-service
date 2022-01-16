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

const message = "this is a test"

func createTestUser() *entity.User {
	return service.CreateDefaultUserService().CreateUser(test.CreateTestUser())
}

func Test_PostService_CreateNewPost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()

	// when
	response, err := postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response != nil)
}

func Test_PostService_CreateNewPost_Fails_WhenUserNotFound(t *testing.T) {
	// setup
	userUuid, _ := uuid.NewRandom()
	postService := service.CreateDefaultPostService()

	// when
	response, err := postService.CreatePost(model.CreateNewPost(&userUuid, message))

	// then
	test.Assert(t, err != nil)
	test.Assert(t, response == nil)
}

func Test_PostService_Can_DeletePost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()
	postModel, _ := postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))

	// when
	err := postService.DeletePost(*postModel.Uuid, *testUser.Uuid)

	// then
	test.Assert(t, err == nil)
}

func Test_PostService_CannotGet_DeletedPost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()
	postModel, _ := postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))
	_ = postService.DeletePost(*postModel.Uuid, *testUser.Uuid)

	// when
	response, err := postService.GetPost(*postModel.Uuid)

	// then
	test.Assert(t, err != nil)
	test.Assert(t, response == nil)
}

func Test_GetAllPosts(t *testing.T) {
	// setup
	postService := service.CreateDefaultPostService()

	// when
	posts := postService.GetAllPosts(constants.UserPostsDefaultPageSize)

	test.Assert(t, posts != nil)
}

func Test_GetPost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()

	// given
	post, err := postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))

	// expect
	test.Assert(t, post != nil)
	test.Assert(t, err == nil)

	// when
	response, err := postService.GetPost(*post.Uuid)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response != nil && response.Text == message)
}

func Test_GetPost_Fails_WhenNotFound(t *testing.T) {
	// setup
	postService := service.CreateDefaultPostService()

	// when
	post, err := postService.GetPost(uuid.New())

	// then
	test.Assert(t, err != nil)
	test.Assert(t, post == nil)
}

func Test_GetPost_Fails_WhenUser_IsNotFound(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()
	post, err := postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))

	// given
	_ = service.CreateDefaultUserService().DeleteUser(*testUser.Uuid)

	// when
	response, err := postService.GetPost(*post.Uuid)

	// then
	test.Assert(t, err != nil)
	test.Assert(t, response == nil)
	test.Assert(t, err.Error() == constants.ErrorMessageUserNotFound)
}

func Test_PostService_GetUserPosts(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()

	// given
	for i := 0; i < 5; i++ {
		_, _ = postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))
	}

	// when
	posts, _ := postService.GetPostsForUser(*testUser.Uuid, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, len(posts) == 5)
}

func Test_PostService_GetUserPosts_FailsFor_MissingUser(t *testing.T) {
	// setup
	testUserUuid, _ := uuid.NewRandom()
	postService := service.CreateDefaultPostService()

	// given
	for i := 0; i < 5; i++ {
		_, _ = postService.CreatePost(model.CreateNewPost(&testUserUuid, message))
	}

	// when
	posts, err := postService.GetPostsForUser(testUserUuid, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, posts == nil)
	test.Assert(t, err != nil)
}

func Test_GetPostsForUserFollows_HappyPath(t *testing.T) {
	// setup
	testUser := createTestUser()
	follower := createTestUser()
	postService := service.CreateDefaultPostService()
	followService := service.CreateDefaultFollowService()
	_, _ = followService.CreateFollow(*follower.Uuid, &model.NewFollow{Following: model.User{ Uuid: testUser.Uuid.String() } })

	// given
	for i := 0; i < 5; i++ {
		_, _ = postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))
	}

	// when
	posts, err := postService.GetPostsForUserFollows(*follower.Uuid, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, len(posts) == 5)
}
