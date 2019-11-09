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

type Post struct {
	Uuid *uuid.UUID `json:"uuid"`

	Message Message `json:"message"`

	Replies []Reply `json:"replies"`

	CreatedAt time.Time `json:"created_at"`
}
