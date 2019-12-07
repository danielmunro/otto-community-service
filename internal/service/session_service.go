package service

import "github.com/danielmunro/otto-community-service/internal/auth/model"

type SessionService struct {
	authService *AuthService
}

func CreateDefaultSessionService() *SessionService {
	return &SessionService{ authService: CreateDefaultAuthService() }
}

func (s *SessionService) CreateSession(newSession *model.NewSession) (*model.Session, error) {
	return s.authService.CreateSession(*newSession)
}
