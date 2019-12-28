package kafka

import (
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/segmentio/kafka-go"
)

func GetReader() *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     string(constants.Users),
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	return r
}
