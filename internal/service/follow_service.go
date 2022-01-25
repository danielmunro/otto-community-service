package service

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/google/uuid"
	"log"
	"time"
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

func (f *FollowService) GetUserFollowers(username string) ([]*model.Follow, error) {
	user, err := f.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}

	follows := f.followRepository.FindByFollowing(user)
	return mapper.GetFollowsModelFromEntities(follows), err
}

func (f *FollowService) GetUserFollowersByUsername(username string) ([]*model.Follow, error) {
	user, err := f.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}

	follows := f.followRepository.FindByFollowing(user)
	return mapper.GetFollowsModelFromEntities(follows), err
}

func (f *FollowService) GetUserFollows(username string) ([]*model.Follow, error) {
	user, err := f.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}

	follows := f.followRepository.FindByUser(user)
	return mapper.GetFollowsModelFromEntities(follows), err
}

func (f *FollowService) GetUserFollowsByUsername(username string) ([]*model.Follow, error) {
	user, err := f.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}

	follows := f.followRepository.FindByUser(user)
	return mapper.GetFollowsModelFromEntities(follows), err
}

func (f *FollowService) DeleteFollow(followUuid uuid.UUID, userUuid uuid.UUID) error {
	follow := f.followRepository.FindOne(followUuid)
	if follow == nil {
		log.Print("follow not found :: ", followUuid)
		return errors.New("follow not found")
	}
	user, _ := f.userRepository.FindOneByUuid(userUuid.String())
	if follow.UserID != user.ID {
		return errors.New("not allowed")
	}
	deletedAt := time.Now()
	follow.DeletedAt = &deletedAt
	f.followRepository.Update(follow)
	return nil
}
