package service

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/google/uuid"
)

type FollowService struct {
	userRepository   *repository.UserRepository
	followRepository *repository.FollowRepository
}

func CreateDefaultFollowService() *FollowService {
	conn := db.CreateDefaultConnection()
	return CreateFollowService(
		repository.CreateUserRepository(conn),
		repository.CreateFollowRepository(conn))
}

func CreateFollowService(userRepository *repository.UserRepository, followRepository *repository.FollowRepository) *FollowService {
	return &FollowService{
		userRepository,
		followRepository,
	}
}

func (f *FollowService) CreateFollow(userUuid uuid.UUID, follow *model.NewFollow) (*model.Follow, error) {
	user, err := f.userRepository.FindOneByUuid(userUuid.String())
	if err != nil {
		return nil, err
	}

	followingUser, err := f.userRepository.FindOneByUuid(follow.Following.Uuid)
	if err != nil {
		return nil, errors.New("follower not found")
	}

	followEntity := entity.GetFollowEntityFromModel(user, followingUser)
	f.followRepository.Create(followEntity)
	return mapper.GetFollowModelFromEntity(followEntity, user, followingUser), nil
}

func (f *FollowService) GetUserFollowers(userUuid uuid.UUID) ([]*model.Follow, error) {
	user, err := f.userRepository.FindOneByUuid(userUuid.String())
	if err != nil {
		return nil, err
	}

	follows := f.followRepository.FindByUser(user)
	return mapper.GetFollowsModelFromEntities(follows, user), err
}
