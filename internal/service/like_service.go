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
	}
}

func CreateLikeService(likeRepository *repository.LikeRepository) *LikeService {
	return &LikeService{
		likeRepository: likeRepository,
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
	postLike := &entity.PostLike{
		Post: post,
		User: user,
	}
	l.likeRepository.CreatePostLike(postLike)
	return mapper.GetPostLikeModelFromEntity(postLike), nil
}
