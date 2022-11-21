package service

import (
	model2 "github.com/danielmunro/otto-community-service/internal/auth/model"
	"github.com/danielmunro/otto-community-service/internal/model"
)

type SecurityService struct{}

func CreateSecurityService() *SecurityService {
	return &SecurityService{}
}

func (s *SecurityService) CanCreateNewPost(session *model2.Session, newPost *model.NewPost) bool {
	return session != nil && session.User.Uuid == newPost.User.Uuid
}
