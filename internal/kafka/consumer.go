package kafka

import (
	"context"
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/segmentio/kafka-go"
	"log"
)


func InitializeAndRunLoop(kafkaHost string) {
	reader := GetReader(kafkaHost)
	userRepository := repository.CreateUserRepository(db.CreateDefaultConnection())
	err := loopKafkaReader(userRepository, reader)
	if err != nil {
		log.Fatal(err)
	}
}

func loopKafkaReader(userRepository *repository.UserRepository, reader *kafka.Reader) error {
	for {
		data, err := reader.ReadMessage(context.Background())
		if err != nil  {
			log.Print(err)
			return nil
		}
		userModel, err := model.DecodeMessageToUser(data.Value)
		if err != nil {
			log.Print("error decoding message to user, skipping", string(data.Value))
			continue
		}
		userEntity, err := userRepository.FindOneByUuid(userModel.Uuid)
		if err == nil {
			userEntity.UpdateUserProfileFromModel(userModel)
			userRepository.Update(userEntity)
		} else {
			userEntity = mapper.GetUserEntityFromModel(userModel)
			userRepository.Create(userEntity)
		}
	}
}
