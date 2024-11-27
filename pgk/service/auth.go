package service

import (
	"rest_api_learn/models"
	"rest_api_learn/pgk/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService (repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}


func (s *AuthService) CreateUser(user models.User) (int,error) {
	return s.repo.CreateUser(user)
}