/*
 * Otto user service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"github.com/danielmunro/otto-community-service/internal"
	"github.com/danielmunro/otto-community-service/internal/middleware"
	"log"
	"net/http"
)

func main() {
	router := internal.NewRouter()
	log.Print("listening on 8081")
	log.Fatal(http.ListenAndServe(":8081",
		middleware.ContentTypeMiddleware(router)))
}
