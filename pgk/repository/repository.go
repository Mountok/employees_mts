package repository

import (
	"rest_api_learn/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(models.User) (int,error)
}

type Employees interface {
	ReadEmployer(models.EmployersResponse) ([]models.EmployersResponse,error)
	ReadEmployers() ([][]models.EmployersResponse, error)
}

type Filters interface {
	ReadAllFiltersDate() (models.FiltersResponse,error)
}

type Repository struct {
	Authorization
	Employees
	Filters
}

func NewRepository(db *sqlx.DB) *Repository { return &Repository{
	Authorization: NewAuthPostgres(db),
	Employees: NewEmployeesPostgres(db),
	Filters: NewFilterPostgres(db),
} }
