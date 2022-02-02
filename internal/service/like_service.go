package service

import (
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/google/uuid"
)

type LikeService struct {
	likeRepository *repository.LikeRepository
	postRepository *repository.PostRepository
	userRepository *repository.UserRepository
}

func CreateDefaultLikeService() *LikeService {
	conn := db.CreateDefaultConnection()
	return &LikeService{
		likeRepository: repository.CreateLikeRepository(conn),
		postRepository: repository.CreatePostRepository(conn),
		userRepository: repository.CreateUserRepository(conn),
	}
}

func (l *LikeService) CreateLikeForPost(postUuid uuid.UUID, userUuid uuid.UUID) (*model.PostLike, error) {
	post, err := l.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		return nil, err
	}
	user, err := l.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return nil, err
	}
	newPostLike := &entity.PostLike{
		Post: post,
		User: user,
	}
	l.likeRepository.Create(newPostLike)
	return mapper.GetPostLikeModelFromEntity(newPostLike), nil
}

func (l *LikeService) DeleteLikeForPost(postUuid uuid.UUID, userUuid uuid.UUID) error {
	post, err := l.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		return err
	}
	user, err := l.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return err
	}
	postLike, err := l.likeRepository.FindByPostAndUser(post, user)
	if err != nil {
		return nil
	}
	l.likeRepository.DeletePostLike(postLike)
	return nil
}
