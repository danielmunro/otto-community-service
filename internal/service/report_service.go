package service

import (
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
)

type ReportService struct {
	userRepository *repository.UserRepository
	postRepository *repository.PostRepository
	replyRepository *repository.ReplyRepository
	reportRepository *repository.ReportRepository
}

func CreateDefaultReportService() *ReportService {
	conn := db.CreateDefaultConnection()
	return CreateReportService(
		repository.CreateUserRepository(conn),
		repository.CreatePostRepository(conn),
		repository.CreateReplyRepository(conn),
		repository.CreateReportRepository(conn))
}

func CreateReportService(
	userRepository *repository.UserRepository,
	postRepository *repository.PostRepository,
	replyRepository *repository.ReplyRepository,
	reportRepository *repository.ReportRepository) *ReportService {
	return &ReportService{
		userRepository,
		postRepository,
		replyRepository,
		reportRepository,
	}
}

func (r *ReportService) CreatePostReport(newReport *model.NewPostReport) (*model.PostReport, error) {
	user, err := r.userRepository.FindOneByUuid(newReport.Message.User.Uuid)
	if err != nil {
		return nil, err
	}

	post, err := r.postRepository.FindOneByUuid(*newReport.Post.Uuid)
	if err != nil {
		return nil, err
	}

	report := entity.CreateReportPostEntity(user, post, newReport)
	r.reportRepository.Create(report)

	return mapper.GetPostReportModelFromEntity(user, post, report), nil
}

func (r *ReportService) CreateReplyReport(newReport *model.NewReplyReport) (*model.ReplyReport, error) {
	user, err := r.userRepository.FindOneByUuid(newReport.Message.User.Uuid)
	if err != nil {
		return nil, err
	}

	reply, err := r.replyRepository.FindOneByUuid(*newReport.Reply.Uuid)
	if err != nil {
		return nil, err
	}

	report := entity.CreateReportReplyEntity(user, reply, newReport)
	r.reportRepository.Create(report)

	return mapper.GetReplyReportModelFromEntity(user, reply, report), nil
}
