package mapper

import (
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/model"
)

func GetPostReportModelFromEntity(user *entity.User, post *entity.Post, report *entity.Report) *model.PostReport {
	return &model.PostReport{
		Uuid:      report.Uuid.String(),
		CreatedAt: report.CreatedAt,
		Text:      report.Text,
		User:      model.User{Uuid: user.Uuid.String()},
		Post: model.Post{
			Uuid: post.Uuid.String(),
		},
	}
}

func GetReplyReportModelFromEntity(user *entity.User, reply *entity.Reply, report *entity.Report) *model.ReplyReport {
	return &model.ReplyReport{
		Uuid:      report.Uuid.String(),
		CreatedAt: report.CreatedAt,
		Text:      report.Text,
		User:      model.User{Uuid: user.Uuid.String()},
		Reply: model.Reply{
			Uuid: reply.Uuid.String(),
		},
	}
}
