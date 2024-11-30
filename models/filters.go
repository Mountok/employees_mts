package models



type DepartmentsModels struct {
	Id int `json:"values" db:"id"`
	Department string `json:"label" db:"department_name"`
}
type JobModels struct {
	Id int `json:"values" db:"id"`
	Job string `json:"label" db:"job_name"`
}
type OfficeModels struct {
	Id int `json:"values" db:"id"`
	Office string `json:"label" db:"office_name"`
}
type RoleModels struct {
	Id int `json:"values" db:"id"`
	Role string `json:"label" db:"role_name"`
}
type SubDivisionModels struct {
	Id int `json:"values" db:"id"`
	SubDivision string `json:"label" db:"subdivision_name"`
}
type BlockModels struct {
	Id int `json:"values" db:"id"`
	Block string `json:"label" db:"block_name"`
}

type FiltersResponse struct {
	Departments []DepartmentsModels `json:"departments"`
	Jobs []JobModels `json:"jobs"`
	Offices []OfficeModels `json:"offices"`
	Role []RoleModels `json:"roles"`
	SubDivision []SubDivisionModels `json:"subdivisions"`
	Block []BlockModels `json:"blocks"`
}