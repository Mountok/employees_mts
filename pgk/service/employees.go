package service

import (
	"rest_api_learn/models"
	"rest_api_learn/pgk/repository"
)

type EmployeesService struct {
	repo repository.Employees
}

func NewEmployeesService(repo repository.Employees) *EmployeesService {
	return &EmployeesService{
		repo: repo,
	}
}


func (s *EmployeesService) ReadEmployers() ([][]models.EmployersResponse, error) {
	
	return s.repo.ReadEmployers()
}


func (s *EmployeesService) ReadEmployer(input models.Employers) ([]models.EmployersResponse, error) {
	return s.repo.ReadEmployer(models.EmployersResponse{
		Id:            input.Id,
		FullName:      input.FullName,
		Number:        input.Number,
		Address:       input.Address,
		City:          input.City,
		JobId:         input.JobId,
		RoleId:        input.RoleId,
		ParentId:      input.ParentId,
		DepartmentId:  input.DepartmentId,
		BlockId:       input.BlockId,
		SubDivisionId: input.SubDivisionId,
	})
}
 