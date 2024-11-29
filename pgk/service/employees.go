package service

import (
	"encoding/json"
	"rest_api_learn/models"
	"rest_api_learn/pgk/repository"

	"github.com/sirupsen/logrus"
)

type EmployeesService struct {
	repo repository.Employees
}

func NewEmployeesService(repo repository.Employees) *EmployeesService {
	return &EmployeesService{
		repo: repo,
	}
}

func (s *EmployeesService) ReadEmployer(input models.Employers) ([]models.EmployersResponse, error) {
	var attributesJSON map[string]string
	err := json.Unmarshal([]byte(input.Attributes), &attributesJSON)
	if err != nil {
		logrus.Errorf("error unmarsjaling json: %s", err)
	}
	logrus.Info(attributesJSON)

	return s.repo.ReadEmployer(models.EmployersResponse{
		Id:           input.Id,
		Name:         input.Name,
		Attributes:   attributesJSON,
		Position:     input.Position,
		DepartmentId: input.DepartmentId,
		ManagerId:    input.ManagerId,
	})
}
