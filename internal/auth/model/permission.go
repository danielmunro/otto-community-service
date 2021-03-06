/*
 * Otto user service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model
// Permission the model 'Permission'
type Permission string

// List of Permission
const (
	READ Permission = "read"
	WRITE Permission = "write"
	MODERATE Permission = "moderate"
	ADMIN Permission = "admin"
)
