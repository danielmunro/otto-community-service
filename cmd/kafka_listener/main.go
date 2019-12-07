package main

import (
	"context"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	_ = godotenv.Load()
	log.Print("connecting to localhost:9092")
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", string(constants.Users), 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("reading from kafka")
	_ = conn.SetReadDeadline(time.Now().Add(10*time.Second))
	for {
		batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
		log.Print("sanity")
		userRepository := repository.CreateUserRepository(db.CreateDefaultConnection())
		log.Print("starting batch read")
		err := ParseBatch(userRepository, batch)
		if err != nil {
			break
		}
		_ = batch.Close()
	}
	_ = conn.Close()
}

func ParseBatch(userRepository *repository.UserRepository, batch *kafka.Batch) error {
	b := make([]byte, 10e3)            // 10KB max per message
	for {
		readLen, err := batch.Read(b)
		if err != nil && err.Error() == "EOF" {
			break
		}
		if err != nil {
			log.Print("error received", err)
			return err
		}
		data := b[:readLen]
		log.Print("received user", string(data))
		userModel := model.DecodeMessageToUser(data)
		userEntity := mapper.GetUserEntityFromModel(userModel)
		userRepository.Create(userEntity)
	}
	return nil
}
