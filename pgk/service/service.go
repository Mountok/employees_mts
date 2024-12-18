package service

import (
	"rest_api_learn/models"
	"rest_api_learn/pgk/repository"
)

type Authorization interface{
	CreateUser(models.User) (int,error)
}

type Employees interface{
	ReadEmployer(models.Employers) ([]models.EmployersResponse,error)
	ReadEmployers() ([][]models.EmployersResponse, error)
}

type Filters interface{
	ReadAllFiltersDate() (models.FiltersResponse,error)
}

type Service struct {
	Authorization
	Employees
	Filters
}

func NewService(repos *repository.Repository) *Service {
	
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Employees: NewEmployeesService(repos.Employees),
		Filters: NewFilterService(repos.Filters),
	}
}