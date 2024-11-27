package models

type EmployersMapResponse struct {
	EmployersResponse
}


type EmployersResponse struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Attributes   int    `json:"attributes" db:"extra_info"`
	Position     string `json:"position" db:"position"`
	DepartmentId int    `json:"department_id" db:"department_id"`
	ManagerId    int	`json:"manager_id" db:"manager_id"`
	Children []Employers
}



type Employers struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Attributes   int    `json:"attributes" db:"extra_info"`
	Position     string `json:"position" db:"position"`
	DepartmentId int    `json:"department_id" db:"department_id"`
	ManagerId    int	`json:"manager_id" db:"manager_id"`

}
