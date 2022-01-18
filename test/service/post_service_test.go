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

func Test_PostService_CreatePublic_NewPost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()

	// when
	response, err := postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response != nil)
	test.Assert(t, response.Visibility == model.PUBLIC)
}

func Test_PostService_CreateNewPost_WithPrivateVisibility(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()

	// given
	newPost := model.CreateNewPost(testUser.Uuid, message)
	newPost.Visibility = model.PRIVATE

	// when
	response, err := postService.CreatePost(newPost)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response.Visibility == model.PRIVATE)
}

func Test_PostService_Respects_PrivateVisibility(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()
	newPost := model.CreateNewPost(testUser.Uuid, message)
	newPost.Visibility = model.PRIVATE
	response, err := postService.CreatePost(newPost)

	// when
	post, err := postService.GetPost(nil, *response.Uuid)

	// then
	test.Assert(t, post == nil)
	test.Assert(t, err != nil)
}

func Test_PostService_CreateNewPost_WithFollowingVisibility(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()

	// given
	newPost := model.CreateNewPost(testUser.Uuid, message)
	newPost.Visibility = model.FOLLOWING

	// when
	response, err := postService.CreatePost(newPost)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response.Visibility == model.FOLLOWING)
}

func Test_PostService_Respects_FollowingVisibility(t *testing.T) {
	// setup
	testUser1 := createTestUser()
	testUser2 := createTestUser()
	testUser3 := createTestUser()
	_, _ = service.CreateDefaultFollowService().CreateFollow(
		*testUser1.Uuid, &model.NewFollow{Following: model.User{Uuid: testUser2.Uuid.String()}})
	postService := service.CreateDefaultPostService()
	newPost := model.CreateNewPost(testUser1.Uuid, message)
	newPost.Visibility = model.FOLLOWING
	response, _ := postService.CreatePost(newPost)

	// when
	post1, err1 := postService.GetPost(testUser2.Uuid, *response.Uuid)
	post2, err2 := postService.GetPost(testUser3.Uuid, *response.Uuid)

	// then
	test.Assert(t, post1 != nil)
	test.Assert(t, err1 == nil)

	test.Assert(t, post2 == nil)
	test.Assert(t, err2 != nil)
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
	response, err := postService.GetPost(nil, *postModel.Uuid)

	// then
	test.Assert(t, err != nil)
	test.Assert(t, response == nil)
}

func Test_GetPosts(t *testing.T) {
	// setup
	postService := service.CreateDefaultPostService()
	testUser := createTestUser()

	// when
	posts, err := postService.GetPosts(testUser.Uuid, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, posts != nil)
}

func Test_GetPosts_NoSession(t *testing.T) {
	// setup
	postService := service.CreateDefaultPostService()

	// when
	posts, err := postService.GetPosts(nil, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, err == nil)
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
	response, err := postService.GetPost(nil, *post.Uuid)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response != nil && response.Text == message)
}

func Test_GetPost_Fails_WhenNotFound(t *testing.T) {
	// setup
	postService := service.CreateDefaultPostService()

	// when
	post, err := postService.GetPost(nil, uuid.New())

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
	response, err := postService.GetPost(nil, *post.Uuid)

	// then
	test.Assert(t, err != nil)
	test.Assert(t, response == nil)
	test.Assert(t, err.Error() == constants.ErrorMessageUserNotFound)
}

func Test_GetPostForUserFollows_Fails_WhenUser_IsNotFound(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreateDefaultPostService()
	_, _ = postService.CreatePost(model.CreateNewPost(testUser.Uuid, message))

	// given
	_ = service.CreateDefaultUserService().DeleteUser(*testUser.Uuid)

	// when
	response, err := postService.GetPostsForUserFollows(*testUser.Uuid, constants.UserPostsDefaultPageSize)

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

func Test_CanGetPosts_ForUserFollows(t *testing.T) {
	// setup
	testUser := createTestUser()
	following := createTestUser()
	postService := service.CreateDefaultPostService()
	followService := service.CreateDefaultFollowService()
	_, _ = followService.CreateFollow(*testUser.Uuid, &model.NewFollow{Following: model.User{Uuid: following.Uuid.String()}})

	// given
	for i := 0; i < 5; i++ {
		_, _ = postService.CreatePost(model.CreateNewPost(following.Uuid, message))
	}

	// when
	posts, err := postService.GetPostsForUserFollows(*testUser.Uuid, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, len(posts) == 5)
}
