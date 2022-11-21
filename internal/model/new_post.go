/*
 * Otto user service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type NewPost struct {
	Text       string     `json:"text"`
	Draft      bool       `json:"draft"`
	User       User       `json:"user"`
	Visibility Visibility `json:"access,omitempty"`
	Images     []NewImage `json:"images,omitempty"`
}

func (n *NewPost) GetOwnerUUID() string {
	return n.User.Uuid
}

func DecodeRequestToNewPost(r *http.Request) (*NewPost, error) {
	decoder := json.NewDecoder(r.Body)
	var data *NewPost
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecodeRequestToPost(r *http.Request) (*Post, error) {
	decoder := json.NewDecoder(r.Body)
	var data *Post
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CreateNewPost(userUuid *uuid.UUID, message string) *NewPost {
	return &NewPost{
		Text: message,
		User: User{
			Uuid: userUuid.String(),
		},
	}
}
