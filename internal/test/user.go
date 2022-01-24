package test

import (
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func CreateTestUser() *model.User {
	randomInt := rand.Int()
	return &model.User{
		Uuid:     uuid.New().String(),
		Username: "user" + strconv.Itoa(randomInt),
	}
}
