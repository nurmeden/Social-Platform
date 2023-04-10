package services

import (
	"social-forum/api/repository"
	"social-forum/models"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) error {
	// user.Password = generatePasswordHash(user.Password)
	s.repo.CreateUser(user)
	return nil
}
