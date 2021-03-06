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
	"net/http"
)

type NewPostReport struct {
	Text string `json:"text,omitempty"`

	User User `json:"user"`

	Post Post `json:"post"`
}

func DecodeRequestToNewPostReport(r *http.Request) *NewPostReport {
	decoder := json.NewDecoder(r.Body)
	var data *NewPostReport
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func DecodeRequestToNewReplyReport(r *http.Request) *NewReplyReport {
	decoder := json.NewDecoder(r.Body)
	var data *NewReplyReport
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}
