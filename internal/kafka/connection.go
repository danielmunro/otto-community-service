package kafka

import (
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"log"
	"os"
	"time"
)

func GetReader() *kafka.Reader {
	mechanism, err := scram.Mechanism(
		scram.SHA512,
		os.Getenv("KAFKA_SASL_USERNAME"),
		os.Getenv("KAFKA_SASL_PASSWORD"))
	if err != nil {
		log.Panic(err)
	}
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_BOOTSTRAP_SERVERS")},
		Topic:     string(constants.Users),
		GroupID: "user_service",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		Dialer: dialer,
	})
	return r
}
