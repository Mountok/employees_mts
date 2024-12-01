package repository

import (
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

func (db *EmployeesPostgres) ReadEmployers() ([][]models.EmployersResponse, error) {

	var idxs []int

	query := "SELECT DISTINCT emp.parent_id FROM main.employees emp WHERE parent_id IS NOT NULL ORDER BY parent_id"

	db.db.Select(&idxs,query)

	fmt.Println(idxs)

	return [][]models.EmployersResponse{}, nil


}
 

func (db *EmployeesPostgres) ReadEmployer(input models.EmployersResponse) ([]models.EmployersResponse, error) {
	var output []models.Employers
	var outputArray []models.EmployersResponse
	var queryString string

	if input.FullName != "" && input.FullName != " " {
		queryString += fmt.Sprint(" and LOWER(emp.full_name) LIKE LOWER('%"+ input.FullName+"%')")
	}
	if input.DepartmentId != "" {
		queryString += fmt.Sprintf(" and emp.departments_id = %s", input.DepartmentId)
	}
	if input.JobId != "" {
		queryString += fmt.Sprintf(" and emp.job_title_id = %s", input.JobId)
	}
	if input.RoleId != "" {
		queryString += fmt.Sprintf(" and emp.role_id = %s", input.RoleId)
	}
	if input.BlockId != "" {
		queryString += fmt.Sprintf(" and emp.block_id = %s", input.BlockId)
	}
	if input.SubDivisionId != "" {
		queryString += fmt.Sprintf(" and emp.subdivision_id = %s", input.SubDivisionId)
	}
	if input.Office != "" {
		queryString += fmt.Sprintf(" and emp.office_id = %s", input.Office)
	}

	logrus.Printf(queryString)

	err := db.db.Select(&output, fmt.Sprintf(
		`SELECT emp.id, emp.full_name, emp.number, emp.adres,emp.citi, job.job_name,rl.role_name, emp.parent_id, dep.department_name, blk.block_name, subd.subdivision_name, offic.office_name FROM main.employees emp 
			INNER JOIN main.role rl  ON  rl.id = emp.role_id
				INNER JOIN main.job_title job  ON job.id = emp.job_title_id
					INNER JOIN main.departments dep  ON dep.id = emp.departments_id
						INNER JOIN main.block blk  ON blk.id = emp.block_id
							INNER JOIN main.subdivision subd  ON subd.id = emp.subdivision_id
								INNER JOIN main.office offic  ON offic.id = emp.office_id 
		WHERE 1=1 %s LIMIT 1`,
		queryString))
	if err != nil {
		return []models.EmployersResponse{}, err
	} else {

		outputArray = append(outputArray, models.EmployersResponse{
			Id:            output[0].Id,
			FullName:      output[0].FullName,
			Number:        output[0].Number,
			Address:       output[0].Address,
			City:          output[0].City,
			JobId:         output[0].JobId,
			RoleId:        output[0].RoleId,
			ParentId:      output[0].ParentId,
			DepartmentId:  output[0].DepartmentId,
			BlockId:       output[0].BlockId,
			SubDivisionId: output[0].SubDivisionId,
		})
	}

	var isSearch = true
	if output[0].ParentId != 0 {
		var lastId = output[0].ParentId
		for isSearch {
			output = []models.Employers{}
			err := db.db.Select(&output, fmt.Sprintf(
				`SELECT emp.id, emp.full_name, emp.number, emp.adres, emp.citi, job.job_name,rl.role_name, emp.parent_id, dep.department_name, blk.block_name, subd.subdivision_name, offic.office_name FROM main.employees emp 
						INNER JOIN main.role rl  ON  rl.id = emp.role_id
							INNER JOIN main.job_title job  ON job.id = emp.job_title_id
								INNER JOIN main.departments dep  ON dep.id = emp.departments_id
									INNER JOIN main.block blk  ON blk.id = emp.block_id
										INNER JOIN main.subdivision subd  ON subd.id = emp.subdivision_id
											INNER JOIN main.office offic  ON offic.id = emp.office_id
					WHERE emp.id=%d`,
				lastId))
			if err != nil {
				return []models.EmployersResponse{}, err
			}
			if output[0].ParentId == 0 {
				isSearch = false
			}

			outputArray = append(outputArray, models.EmployersResponse{
				Id:            output[0].Id,
			FullName:      output[0].FullName,
			Number:        output[0].Number,
			Address:       output[0].Address,
			City:          output[0].City,
			JobId:         output[0].JobId,
			RoleId:        output[0].RoleId,
			ParentId:      output[0].ParentId,
			DepartmentId:  output[0].DepartmentId,
			BlockId:       output[0].BlockId,
			SubDivisionId: output[0].SubDivisionId,
			})
			fmt.Println(outputArray)
			lastId = output[0].ParentId
		}
	}

	slices.Reverse(outputArray)

	return outputArray, nil
}

// func (db *EmployeesPostgres) ReadEmployer(input models.EmployersResponse) ([]models.EmployersResponse, error) {
// 	var output []models.Employers
// 	var outputArray []models.EmployersResponse
// 	var queryString string
// 	for key := range input.Attributes {
// 		queryString += fmt.Sprintf("and extra_info ->> '%s' = '%s' ", key, input.Attributes[key])
// 	}

// 	if input.DepartmentId != 0 {
// 		queryString += fmt.Sprintf("and emp.department_id=%d ", input.DepartmentId)
// 	}
// 	if input.ManagerId != 0 {
// 		queryString += fmt.Sprintf("and emp.manager_id=%d ", input.ManagerId)
// 	}
// 	if input.Name != "" {
// 		queryString += fmt.Sprintf("and emp.name LIKE '%s%s%' ", "%", input.Name)
// 	}
// 	if input.Position != "" {
// 		queryString += fmt.Sprintf("and emp.position='%s' ", "%", input.Position)
// 	}

// 	logrus.Printf(queryString)

// 	err := db.db.Select(&output, fmt.Sprintf(
// 		"SELECT emp.id,emp.name, ps.position, emp.department_id, emp.manager_id, emp.extra_info  FROM employees emp JOIN positions ps ON ps.id = emp.position WHERE 1=1 %s LIMIT 1",
// 		queryString))
// 	if err != nil {
// 		return []models.EmployersResponse{}, err
// 	} else {
// 		fmt.Println(output)
// 		var attributesJSON map[string]string
// 		err = json.Unmarshal([]byte(output[0].Attributes), &attributesJSON)
// 		if err != nil {
// 			logrus.Errorf("error unmarsjaling json: %s", err)
// 			return []models.EmployersResponse{}, nil
// 		}
// 		outputArray = append(outputArray, models.EmployersResponse{
// 			Id:           output[0].Id,
// 			Name:         output[0].Name,
// 			Position:     output[0].Position,
// 			DepartmentId: output[0].DepartmentId,
// 			ManagerId:    output[0].ManagerId,
// 			Attributes:   attributesJSON,
// 		})
// 	}

// 	var isSearch = true
// 	if output[0].ManagerId != 0 {
// 		var lastId = output[0].ManagerId
// 		for isSearch {
// 			output = []models.Employers{}
// 			err := db.db.Select(&output, fmt.Sprintf("SELECT emp.id,emp.name, ps.position, emp.department_id, emp.manager_id, emp.extra_info  FROM employees emp JOIN positions ps ON ps.id = emp.position WHERE emp.id = %d", lastId))
// 			if err != nil {
// 				return []models.EmployersResponse{}, err
// 			}
// 			if output[0].ManagerId == 0 {
// 				isSearch = false
// 			}
// 			var attributesJSON map[string]string
// 			err = json.Unmarshal([]byte(output[0].Attributes), &attributesJSON)
// 			if err != nil {
// 				logrus.Errorf("error unmarsjaling json: %s", err)
// 				return []models.EmployersResponse{}, nil
// 			}
// 			outputArray = append(outputArray, models.EmployersResponse{
// 				Id:           output[0].Id,
// 				Name:         output[0].Name,
// 				Position:     output[0].Position,
// 				DepartmentId: output[0].DepartmentId,
// 				ManagerId:    output[0].ManagerId,
// 				Attributes:   attributesJSON,
// 			})
// 			fmt.Println(outputArray)
// 			lastId = output[0].ManagerId
// 		}
// 	}

// 	slices.Reverse(outputArray)

// 	return outputArray, nil
// }
