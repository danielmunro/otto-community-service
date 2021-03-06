/*
 * Otto user service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package internal

import (
	"github.com/danielmunro/otto-community-service/internal/auth/model"
	"github.com/danielmunro/otto-community-service/internal/controller"
	"github.com/danielmunro/otto-community-service/internal/service"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello world"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("login"))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	res, err := service.CreateDefaultSessionService().CreateSession(
		model.CreateNewSession("dan+6211@danmunro.com", "my-awesome-new-pAssword-123!"))
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write([]byte(res.Token))
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"Login",
		"GET",
		"/login",
		Login,
	},

	{
		"Signup",
		"GET",
		"/signup",
		Signup,
	},

	{
		"AddUserFollowV1",
		strings.ToUpper("Post"),
		"/user/{uuid}/follows",
		controller.CreateNewFollowV1,
	},

	{
		"CreateAReplyV1",
		strings.ToUpper("Post"),
		"/reply",
		controller.CreateAReplyV1,
	},

	{
		"CreateAReplyReportV1",
		strings.ToUpper("Post"),
		"/report/reply",
		controller.CreateReplyReportV1,
	},

	{
		"CreateAPostReport",
		strings.ToUpper("Post"),
		"/report",
		controller.CreatePostReportV1,
	},

	{
		"CreateAReshare",
		strings.ToUpper("Post"),
		"/reshare",
		CreateAReshare,
	},

	{
		"CreateNewPostV1",
		strings.ToUpper("Post"),
		"/post",
		controller.CreateNewPostV1,
	},

	{
		"CreateSessionV1",
		strings.ToUpper("Post"),
		"/session",
		controller.CreateSessionV1,
	},

	{
		"GetPostV1",
		strings.ToUpper("Get"),
		"/post/{uuid}",
		controller.GetPostV1,
	},

	{
		"GetPostRepliesV1",
		strings.ToUpper("Get"),
		"/post/{uuid}/replies",
		controller.GetPostRepliesV1,
	},

	{
		"GetPosts",
		strings.ToUpper("Get"),
		"/post",
		GetPosts,
	},

	{
		"GetReshare",
		strings.ToUpper("Get"),
		"/rehare/{uuid}",
		GetReshare,
	},

	{
		"GetSuggestedFollowsForUserV1",
		strings.ToUpper("Get"),
		"/user/{uuid}/suggested-follows",
		controller.GetSuggestedFollowsForUserV1,
	},

	{
		"GetUserV1",
		strings.ToUpper("Get"),
		"/user/{uuid}",
		controller.GetUserV1,
	},

	{
		"GetUserFollowsV1",
		strings.ToUpper("Get"),
		"/user/{uuid}/follows",
		controller.GetUserFollowsV1,
	},

	{
		"GetUserFollowsPostsV1",
		strings.ToUpper("Get"),
		"/post/follows/{uuid}",
		controller.GetUserFollowsPostsV1,
	},

	{
		"GetUserPostsV1",
		strings.ToUpper("Get"),
		"/user/{uuid}/posts",
		controller.GetUserPostsV1,
	},

	{
		"GetNewPostsV1",
		strings.ToUpper("Get"),
		"/user/{uuid}/new-posts",
		controller.GetNewPostsV1,
	},
}
