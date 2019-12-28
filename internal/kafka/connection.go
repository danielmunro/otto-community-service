package kafka

import (
	"context"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/segmentio/kafka-go"
	"log"
)

func GetConnection() *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", string(constants.Users), 0)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
