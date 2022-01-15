package kafka

import (
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/google/uuid"
	"log"
)

func InitializeAndRunLoop() {
	userRepository := repository.CreateUserRepository(db.CreateDefaultConnection())
	err := loopKafkaReader(userRepository)
	if err != nil {
		log.Fatal(err)
	}
}

func loopKafkaReader(userRepository *repository.UserRepository) error {
	reader := GetReader()
	for {
		log.Print("listening for kafka messages")
		data, err := reader.ReadMessage(-1)
		log.Print("message received")
		if err != nil  {
			log.Print(err)
			return nil
		}
		log.Print("consuming user message ", string(data.Value))
		userModel, err := model.DecodeMessageToUser(data.Value)
		if err != nil {
			log.Print("error decoding message to user, skipping", string(data.Value))
			continue
		}
		_, err = uuid.Parse(userModel.Uuid)
		if err != nil {
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
