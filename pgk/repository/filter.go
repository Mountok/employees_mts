package repository

import (
	"rest_api_learn/models"

	"github.com/jmoiron/sqlx"
)

type FilterPostgres struct {
	db *sqlx.DB
}

func NewFilterPostgres(db *sqlx.DB) *FilterPostgres {
	return &FilterPostgres{db: db}
}

func (s *FilterPostgres) ReadAllFiltersDate() (filters models.FiltersResponse,err error) {

	var departmentsQuery string = "SELECT id, department_name FROM main.departments"
	var departmentsModel []models.DepartmentsModels
	err = s.db.Select(&departmentsModel,departmentsQuery)
	if err != nil {return filters, err}


	var jobQuery string = "SELECT id, job_name FROM main.job_title"
	var jobModels []models.JobModels
	err = s.db.Select(&jobModels,jobQuery)
	if err != nil {return filters, err}

	var officeQuery string = "SELECT id, office_name FROM main.office"
	var officeModels []models.OfficeModels
	err = s.db.Select(&officeModels,officeQuery)
	if err != nil {return filters, err}


	var roleQuery string = "SELECT id, role_name FROM main.role"
	var roleModels []models.RoleModels
	err = s.db.Select(&roleModels,roleQuery)
	if err != nil {return filters, err}

	var subDivisionQuery string = "SELECT id, subdivision_name FROM main.subdivision"
	var subDivisionModels []models.SubDivisionModels
	err = s.db.Select(&subDivisionModels,subDivisionQuery)
	if err != nil {return filters, err}

	var blockQuery string = "SELECT id, block_name FROM main.block"
	var blockModels []models.BlockModels
	err = s.db.Select(&blockModels,blockQuery)
	if err != nil {return filters, err}



	return models.FiltersResponse{
		 Departments: departmentsModel,
		 Jobs: jobModels,
		 Offices: officeModels,
		 Role: roleModels,
		 SubDivision: subDivisionModels,
		 Block: blockModels,
	}, nil
}