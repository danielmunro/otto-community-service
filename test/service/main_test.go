package service_test

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("test main")
	if os.Getenv("CI") == "" {
		_ = godotenv.Load()
	}
	os.Exit(m.Run())
}
