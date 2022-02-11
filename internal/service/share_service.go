package service

import (
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/google/uuid"
)

type ShareService struct {
	shareRepository *repository.ShareRepository
	postRepository  *repository.PostRepository
	userRepository  *repository.UserRepository
}

func CreateDefaultShareService() *ShareService {
	conn := db.CreateDefaultConnection()
	return CreateShareService(
		repository.CreateShareRepository(conn),
		repository.CreatePostRepository(conn),
		repository.CreateUserRepository(conn),
	)
}

func CreateShareService(
	shareRepository *repository.ShareRepository,
	postRepository *repository.PostRepository,
	userRepository *repository.UserRepository) *ShareService {
	return &ShareService{
		shareRepository,
		postRepository,
		userRepository,
	}
}

func (s *ShareService) CreateShare(share *model.NewShare) (*model.Share, error) {
	user, _ := s.userRepository.FindOneByUuid(uuid.MustParse(share.User.Uuid))
	post, _ := s.postRepository.FindOneByUuid(uuid.MustParse(share.Post.Uuid))
	shareEntity := entity.CreateShare(user, post, share)
	s.shareRepository.Save(shareEntity)
	return mapper.GetShareModelFromEntity(shareEntity), nil
}

func (s *ShareService) GetShare(shareUuid uuid.UUID) (*model.Share, error) {
	share, err := s.shareRepository.FindOneByUuid(shareUuid)
	if err != nil {
		return nil, err
	}
	return mapper.GetShareModelFromEntity(share), nil
}
