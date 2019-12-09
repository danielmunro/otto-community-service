package test

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"time"
)

func CreateTestUser() *model.User {
	return &model.User{
		Uuid: uuid.New().String(),
		Name: "tester mctesterson",
		Username: randomdata.SillyName(),
		ProfilePic: "https://icatcare.org/app/uploads/2019/09/The-Kitten-Checklist-1.png",
		Birthday: time.Now().Format("2006-01-02"),
		Location: randomdata.City() + ", " + randomdata.State(randomdata.Large),
		BioMessage: randomdata.Paragraph(),
		Email: randomdata.Email(),
		Phone: randomdata.PhoneNumber(),
	}
}
