package repository

import (
	"encoding/json"
	"fmt"
	"rest_api_learn/models"
	"slices"

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

func (db *EmployeesPostgres) ReadEmployer(input models.EmployersResponse) ([]models.EmployersResponse, error) {
	var output []models.Employers
	var outputArray []models.EmployersResponse
	var queryString string
	for key := range input.Attributes {
		queryString += fmt.Sprintf("and extra_info ->> '%s' = '%s'", key, input.Attributes[key])
	}

	if input.DepartmentId != 0 {
		queryString += fmt.Sprintf("and department_id=%d", input.DepartmentId)
	}

	logrus.Printf(queryString)

	err := db.db.Select(&output, fmt.Sprintf("SELECT id, name, position, department_id, manager_id, extra_info FROM employees WHERE 1=1 %s LIMIT 1", queryString))
	if err != nil {
		return []models.EmployersResponse{}, err
	} else {
		fmt.Println(output)
		var attributesJSON map[string]string
		err = json.Unmarshal([]byte(output[0].Attributes), &attributesJSON)
		if err != nil {
			logrus.Errorf("error unmarsjaling json: %s", err)
			return []models.EmployersResponse{}, nil
		}
		outputArray = append(outputArray, models.EmployersResponse{
			Id:           output[0].Id,
			Name:         output[0].Name,
			Position:     output[0].Position,
			DepartmentId: output[0].DepartmentId,
			ManagerId:    output[0].ManagerId,
			Attributes:   attributesJSON,
		})
	}

	var isSearch = true
	if output[0].ManagerId != 0 {
		var lastId = output[0].ManagerId
		for isSearch {
			output = []models.Employers{}
			err := db.db.Select(&output, fmt.Sprintf("SELECT id, name, position, department_id, manager_id, extra_info FROM employees WHERE id = %d", lastId))
			if err != nil {
				return []models.EmployersResponse{}, err
			}
			if output[0].ManagerId == 0 {
				isSearch = false
			}
			var attributesJSON map[string]string
			err = json.Unmarshal([]byte(output[0].Attributes), &attributesJSON)
			if err != nil {
				logrus.Errorf("error unmarsjaling json: %s", err)
				return []models.EmployersResponse{}, nil
			}
			outputArray = append(outputArray, models.EmployersResponse{
				Id:           output[0].Id,
				Name:         output[0].Name,
				Position:     output[0].Position,
				DepartmentId: output[0].DepartmentId,
				ManagerId:    output[0].ManagerId,
				Attributes:   attributesJSON,
			})
			fmt.Println(outputArray)
			lastId = output[0].ManagerId
		}
	}

	slices.Reverse(outputArray)

	return outputArray, nil
}
