package entity

import (
	"github.com/danielmunro/otto-community-service/internal/enum"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	Text       string
	UserID     uint
	User *User
	Visibility enum.Visibility
	Uuid   *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ReportedID uint
	ReportedType string
}

func CreateReportPostEntity(reporter *User, post *Post, report *model.NewPostReport) *Report {
	return &Report{
		Text:       report.Text,
		UserID:     reporter.ID,
		Visibility: enum.PRIVATE,
		ReportedID: post.ID,
		ReportedType: "Post",
	}
}

func CreateReportReplyEntity(reporter *User, reply *Reply, report *model.NewReplyReport) *Report {
	return &Report{
		Text:       report.Text,
		UserID:     reporter.ID,
		Visibility: enum.PRIVATE,
		ReportedID: reply.ID,
		ReportedType: "Reply",
	}
}
