package repository

import (
	"encoding/json"
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
	var outputArray [][]models.Employers
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
		return models.EmployersResponse{}, err
	} else {
		fmt.Println(output)
		outputArray = append(outputArray, output)
	}

	var isSearch = true
	if output[0].ManagerId != 0 {
		var lastId = output[0].ManagerId
		for isSearch {
			output = []models.Employers{}
			err := db.db.Select(&output, fmt.Sprintf("SELECT id, name, position, department_id, manager_id, extra_info FROM employees WHERE id = %d", lastId))
			if err != nil {
				return models.EmployersResponse{}, err
			}
			if output[0].ManagerId == 0 {
				isSearch = false
			}
			outputArray = append(outputArray, output)
			fmt.Println(outputArray)
			lastId = output[0].ManagerId
		}
	}

	if len(outputArray) == 1 {
		var attributesJSON map[string]string
		err := json.Unmarshal([]byte(outputArray[0][0].Attributes), &attributesJSON)
		if err != nil {
			logrus.Errorf("error unmarsjaling json: %s", err)
			return models.EmployersResponse{}, nil
		}
		return models.EmployersResponse{
			Id:           outputArray[0][0].Id,
			Name:         outputArray[0][0].Name,
			Position:     outputArray[0][0].Position,
			DepartmentId: outputArray[0][0].DepartmentId,
			ManagerId:    outputArray[0][0].ManagerId,
			Attributes:   attributesJSON,
		}, nil
	} else {
		var arrayLen = len(outputArray)
		var attributesJSON map[string]string
		err := json.Unmarshal([]byte(outputArray[0][0].Attributes), &attributesJSON)
		if err != nil {
			logrus.Errorf("error unmarsjaling json: %s", err)
			return models.EmployersResponse{}, nil
		}
		var result = models.EmployersResponse{
			Id:           outputArray[arrayLen-1][0].Id,
			Name:         outputArray[arrayLen-1][0].Name,
			Position:     outputArray[arrayLen-1][0].Position,
			DepartmentId: outputArray[arrayLen-1][0].DepartmentId,
			ManagerId:    outputArray[arrayLen-1][0].ManagerId,
			Attributes:   attributesJSON,
		}

		logrus.Info(result)
		for i := 0; i <= arrayLen-2; i++ {
			var mainBranch = models.EmployersResponse{
				Id:           outputArray[i][0].Id,
				Name:         outputArray[i][0].Name,
				Position:     outputArray[i][0].Position,
				DepartmentId: outputArray[i][0].DepartmentId,
				ManagerId:    outputArray[i][0].ManagerId,
			}
			result.Children = append(result.Children,mainBranch)
			
			fmt.Println(mainBranch)

			fmt.Println(i)
			
		}
		return result, nil
	}

}
