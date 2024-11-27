package service

import (
	"rest_api_learn/models"
	"rest_api_learn/pgk/repository"
)

type Authorization interface{
	CreateUser(models.User) (int,error)
}

type Employees interface{}

type Service struct {
	Authorization
	Employees
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}