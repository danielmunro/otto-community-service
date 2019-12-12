package test

import (
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
)

func CreateTestUser() *model.User {
	return &model.User{
		Uuid: uuid.New().String(),
	}
}
