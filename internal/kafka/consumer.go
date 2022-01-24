package kafka

import (
	"encoding/json"
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
		log.Print("message received on topic :: ", data.TopicPartition.Topic)
		if err != nil {
			log.Print(err)
			return nil
		}
		log.Print("data :: ", data.Value)
		if *data.TopicPartition.Topic == "users" {
			readUser(userRepository, data.Value)
		} else if *data.TopicPartition.Topic == "images" {
			updateUserImage(userRepository, data.Value)
		}
	}
}

func updateUserImage(userRepository *repository.UserRepository, data []byte) {
	result := decodeToMap(data)
	userUuid := result["uuid"].(string)
	s3Key := result["s3_key"].(string)
	log.Print("update user profile pic :: {}, {}, {}", userUuid, s3Key, result)
	userEntity, err := userRepository.FindOneByUuid(userUuid)
	if err != nil {
		log.Print("user not found when updating profile pic")
		return
	}
	log.Print("update user with s3 key", userEntity.Uuid.String(), s3Key)
	userEntity.ProfilePic = s3Key
	userRepository.Update(userEntity)
}

func readUser(userRepository *repository.UserRepository, data []byte) {
	log.Print("consuming user message ", string(data))
	userModel, err := model.DecodeMessageToUser(data)
	if err != nil {
		log.Print("error decoding message to user, skipping", string(data))
		return
	}
	_, err = uuid.Parse(userModel.Uuid)
	if err != nil {
		return
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

func decodeToMap(data []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	return result
}
