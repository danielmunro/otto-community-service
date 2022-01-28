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
	"time"
)

type Image struct {
	Uuid string `json:"uuid"`

	S3Key string `json:"s3_key"`

	CreatedAt time.Time `json:"created_at"`

	Post Post `json:"post"`

	User User `json:"user"`
}
