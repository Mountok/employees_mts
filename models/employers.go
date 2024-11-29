package models


type EmployersResponse struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Attributes   map[string]string `json:"attributes" db:"extra_info"`
	Position     string `json:"position" db:"position"`
	DepartmentId int    `json:"department_id" db:"department_id"`
	ManagerId    int    `json:"manager_id" db:"manager_id"`
	Children     []EmployersResponse `json:"children"`
}

type Employers struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Attributes   string `json:"attributes" db:"extra_info"`
	Position     string `json:"position" db:"position"`
	DepartmentId int    `json:"department_id" db:"department_id"`
	ManagerId    int    `json:"manager_id" db:"manager_id"`
}
