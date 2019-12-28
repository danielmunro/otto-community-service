package kafka

import (
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const SecondsToWait = 20

func Check() {
	conn := GetConnection()
	_ = conn.SetReadDeadline(time.Now().Add(SecondsToWait * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	userRepository := repository.CreateUserRepository(db.CreateDefaultConnection())
	err := ParseBatch(userRepository, batch)
	if err != nil {
		log.Fatal(err)
	}
	_ = batch.Close()
	_ = conn.Close()
}

func ParseBatch(userRepository *repository.UserRepository, batch *kafka.Batch) error {
	b := make([]byte, 10e3) // 10KB max per message
	for {
		readLen, err := batch.Read(b)
		if err != nil  {
			return nil
		}
		data := b[:readLen]
		userModel, err := model.DecodeMessageToUser(data)
		if err != nil {
			log.Print("error decoding message to user, skipping")
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
