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
	"time"
)

type User struct {
	Uuid string `json:"uuid"`

	Username string `json:"username,omitempty"`

	Name string `json:"name,omitempty"`

	Birthday time.Time `json:"birthday,omitempty"`

	BioMessage string `json:"bio_message,omitempty"`

	Role Role `json:"role,omitempty"`

	IsBanned bool `json:"is_banned,omitempty"`

	ProfilePic string `json:"profile_pic,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	Posts []Post `json:"posts,omitempty"`

	Replies []Reply `json:"replies,omitempty"`

	Follows []Follow `json:"follows,omitempty"`
}

func DecodeMessageToUser(message []byte) (*User, error) {
	user := &User{}
	err := json.Unmarshal(message, user)
	return user, err
}
