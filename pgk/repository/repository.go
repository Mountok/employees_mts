package repository

import (
	"rest_api_learn/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(models.User) (int,error)
}

type Employees interface {
	ReadEmployer(models.EmployersResponse) (models.EmployersResponse,error)
}

type Repository struct {
	Authorization
	Employees
}

func NewRepository(db *sqlx.DB) *Repository { return &Repository{
	Authorization: NewAuthPostgres(db),
	Employees: NewEmployeesPostgres(db),
} }
