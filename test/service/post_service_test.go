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
	postService := service.CreatePostService()

	// when
	response, err := postService.CreatePost(*testUser.Uuid, model.CreateNewPost(testUser.Uuid, message))

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response != nil)
	test.Assert(t, response.Visibility == model.PUBLIC)
}

func Test_PostService_CreateNewPost_WithPrivateVisibility(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreatePostService()

	// given
	newPost := model.CreateNewPost(testUser.Uuid, message)
	newPost.Visibility = model.PRIVATE

	// when
	response, err := postService.CreatePost(*testUser.Uuid, newPost)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response.Visibility == model.PRIVATE)
}

func Test_PostService_Respects_PrivateVisibility(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreatePostService()
	newPost := model.CreateNewPost(testUser.Uuid, message)
	newPost.Visibility = model.PRIVATE
	response, err := postService.CreatePost(*testUser.Uuid, newPost)

	// when
	post, err := postService.GetPost(nil, uuid.MustParse(response.Uuid))

	// then
	test.Assert(t, post == nil)
	test.Assert(t, err != nil)
}

func Test_PostService_CreateNewPost_WithFollowingVisibility(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreatePostService()

	// given
	newPost := model.CreateNewPost(testUser.Uuid, message)
	newPost.Visibility = model.FOLLOWING

	// when
	response, err := postService.CreatePost(*testUser.Uuid, newPost)

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
	postService := service.CreatePostService()
	newPost := model.CreateNewPost(testUser1.Uuid, message)
	newPost.Visibility = model.FOLLOWING
	response, _ := postService.CreatePost(*testUser1.Uuid, newPost)

	// when
	post1, err1 := postService.GetPost(testUser2.Uuid, uuid.MustParse(response.Uuid))
	post2, err2 := postService.GetPost(testUser3.Uuid, uuid.MustParse(response.Uuid))

	// then
	test.Assert(t, post1 != nil)
	test.Assert(t, err1 == nil)

	test.Assert(t, post2 == nil)
	test.Assert(t, err2 != nil)
}

func Test_PostService_CreateNewPost_Fails_WhenUserNotFound(t *testing.T) {
	// setup
	userUuid, _ := uuid.NewRandom()
	postService := service.CreatePostService()

	// when
	response, err := postService.CreatePost(userUuid, model.CreateNewPost(&userUuid, message))

	// then
	test.Assert(t, err != nil)
	test.Assert(t, response == nil)
}

func Test_PostService_Can_DeletePost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreatePostService()
	postModel, _ := postService.CreatePost(*testUser.Uuid, model.CreateNewPost(testUser.Uuid, message))

	// when
	err := postService.DeletePost(uuid.MustParse(postModel.Uuid), *testUser.Uuid)

	// then
	test.Assert(t, err == nil)
}

func Test_PostService_CannotGet_DeletedPost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreatePostService()
	postModel, _ := postService.CreatePost(*testUser.Uuid, model.CreateNewPost(testUser.Uuid, message))
	_ = postService.DeletePost(uuid.MustParse(postModel.Uuid), *testUser.Uuid)

	// when
	response, err := postService.GetPost(nil, uuid.MustParse(postModel.Uuid))

	// then
	test.Assert(t, err != nil)
	test.Assert(t, response == nil)
}

func Test_GetPosts(t *testing.T) {
	// setup
	postService := service.CreatePostService()
	testUser := createTestUser()

	// when
	posts, err := postService.GetPostsFirehose(&testUser.Username, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, posts != nil)
}

func Test_GetPosts_NoSession(t *testing.T) {
	// setup
	postService := service.CreatePostService()

	// when
	posts, err := postService.GetPostsFirehose(nil, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, posts != nil)
}

func Test_GetPost(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreatePostService()

	// given
	post, err := postService.CreatePost(*testUser.Uuid, model.CreateNewPost(testUser.Uuid, message))

	// expect
	test.Assert(t, post != nil)
	test.Assert(t, err == nil)

	// when
	response, err := postService.GetPost(nil, uuid.MustParse(post.Uuid))

	// then
	test.Assert(t, err == nil)
	test.Assert(t, response != nil && response.Text == message)
}

func Test_GetPost_Fails_WhenNotFound(t *testing.T) {
	// setup
	postService := service.CreatePostService()

	// when
	post, err := postService.GetPost(nil, uuid.New())

	// then
	test.Assert(t, err != nil)
	test.Assert(t, post == nil)
}

func Test_PostService_GetUserPosts(t *testing.T) {
	// setup
	testUser := createTestUser()
	postService := service.CreatePostService()

	// given
	for i := 0; i < 5; i++ {
		_, _ = postService.CreatePost(*testUser.Uuid, model.CreateNewPost(testUser.Uuid, message))
	}

	// when
	posts, _ := postService.GetPostsForUser(testUser.Username, nil, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, len(posts) == 5)
}

func Test_PostService_GetUserPosts_FailsFor_MissingUser(t *testing.T) {
	// setup
	testUserUuid, _ := uuid.NewRandom()
	postService := service.CreatePostService()

	// given
	for i := 0; i < 5; i++ {
		_, _ = postService.CreatePost(testUserUuid, model.CreateNewPost(&testUserUuid, message))
	}

	// when
	posts, err := postService.GetPostsForUser(testUserUuid.String(), nil, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, posts == nil)
	test.Assert(t, err != nil)
}

func Test_CanGetPosts_ForUserFollows(t *testing.T) {
	// setup
	testUser := createTestUser()
	following := createTestUser()
	postService := service.CreatePostService()
	followService := service.CreateDefaultFollowService()
	_, _ = followService.CreateFollow(*testUser.Uuid, &model.NewFollow{Following: model.User{Uuid: following.Uuid.String()}})

	// given
	for i := 0; i < 5; i++ {
		_, _ = postService.CreatePost(*following.Uuid, model.CreateNewPost(following.Uuid, message))
	}

	// when
	posts, err := postService.GetPostsForUserFollows(testUser.Username, *testUser.Uuid, constants.UserPostsDefaultPageSize)

	// then
	test.Assert(t, err == nil)
	test.Assert(t, len(posts) >= 5)
}
