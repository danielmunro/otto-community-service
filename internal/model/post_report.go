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
	"github.com/google/uuid"
	"time"
)

type PostReport struct {
	Uuid string `json:"uuid"`

	CreatedAt time.Time `json:"created_at"`

	Reporter User `json:"reporter"`

	Message Message `json:"message"`

	Post Post `json:"post"`
}

func CreateNewPostReport(userUuid *uuid.UUID, postUuid *uuid.UUID, message string) *NewPostReport {
	return &NewPostReport{
		Message: NewMessage{
			Text: message,
			User: User{
				Uuid: userUuid.String(),
			},
		},
		Post: Post{
			Uuid: postUuid,
		},
	}
}

func CreateNewReplyReport(userUuid *uuid.UUID, replyUuid *uuid.UUID, message string) *NewReplyReport {
	return &NewReplyReport{
		Message: NewMessage{
			Text: message,
			User: User{
				Uuid: userUuid.String(),
			},
		},
		Reply: Reply{
			Uuid: replyUuid,
		},
	}
}
