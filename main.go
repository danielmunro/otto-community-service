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
	"github.com/danielmunro/otto-community-service/internal/kafka"
	"github.com/danielmunro/otto-community-service/internal/middleware"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
)

func main() {
	go serveHttp()
	readKafka()
}

func readKafka() {
	kafkaHost := os.Getenv("KAFKA_HOST")
	log.Print("connecting to kafka", kafkaHost)
	kafka.InitializeAndRunLoop(kafkaHost)
	log.Print("exit kafka loop")
}

func serveHttp() {
	router := internal.NewRouter()
	log.Print("http listening on 8081")
	log.Fatal(http.ListenAndServe(":8081",
		middleware.ContentTypeMiddleware(router)))
}
