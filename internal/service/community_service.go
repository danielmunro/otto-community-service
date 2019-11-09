package service

import (
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/repository"
)

type CommunityService struct {
	userRepository *repository.UserRepository
	postService    *PostService
}

func CreateDefaultCommunityService() *CommunityService {
	return CreateCommunityService(repository.CreateUserRepository(db.CreateDefaultConnection()))
}

func CreateCommunityService(userRepository *repository.UserRepository) *CommunityService {
	return &CommunityService{
		userRepository: userRepository,
	}
}
