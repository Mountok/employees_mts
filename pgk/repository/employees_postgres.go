package repository

import (
	"fmt"
	"rest_api_learn/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type EmployeesPostgres struct {
	db *sqlx.DB
}

func NewEmployeesPostgres(db *sqlx.DB) *EmployeesPostgres {
	return &EmployeesPostgres{
		db: db,
	}
}

func (db *EmployeesPostgres) ReadEmployer(input models.EmployersResponse) (models.EmployersResponse, error) {
	var output []models.Employers
	var queryString string
	for key := range input.Attributes {
		queryString += fmt.Sprintf("and extra_info ->> '%s' = '%s'", key,input.Attributes[key])
	}

	if input.DepartmentId != 0 {
		queryString += fmt.Sprintf("and department_id=%d",input.DepartmentId)
	}

	logrus.Printf(queryString)

	err := db.db.Select(&output, fmt.Sprintf("SELECT id, name, position, department_id, manager_id, extra_info FROM employees WHERE 1=1 %s",queryString))
	if err != nil {
		return models.EmployersResponse{},err
	}
	return models.EmployersResponse{
		Id: output[0].Id,
		Name: output[0].Name,
		Position: output[0].Position,
		DepartmentId: output[0].DepartmentId,
		ManagerId: output[0].ManagerId,
		Attributes: input.Attributes,
	}, nil
}