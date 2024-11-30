package models

type EmployersResponse struct {
	Id            int                 `json:"id" db:"id"`
	FullName      string              `json:"full_name" db:"full_name"`
	Number        string              `json:"number" db:"number"`
	Address       string              `json:"address" db:"adres"`
	City          string              `json:"city" db:"citi"`
	JobId         string              `json:"job_name" db:"job_name"`
	RoleId        string              `json:"role_name" db:"role_name"`
	ParentId      int                 `json:"parent_id" db:"parent_id"`
	DepartmentId  string              `json:"department_name" db:"department_name"`
	BlockId       string              `json:"block_name db:"block_name"`
	SubDivisionId string              `json:"subdivision_name" db:"subdivision_name"`
	Office        string              `json:"office_name" db:"office_name"`
	Children      []EmployersResponse `json:"children"`
}

type Employers struct {
	Id            int    `json:"id" db:"id"`
	FullName      string `json:"full_name" db:"full_name"`
	Number        string `json:"number" db:"number"`
	Address       string `json:"address" db:"adres"`
	City          string `json:"city" db:"citi"`
	JobId         string `json:"job_name" db:"job_name"`
	RoleId        string `json:"role_name" db:"role_name"`
	ParentId      int    `json:"parent_id" db:"parent_id"`
	DepartmentId  string `json:"department_name" db:"department_name"`
	BlockId       string `json:"block_name" db:"block_name"`
	SubDivisionId string `json:"subdivision_name" db:"subdivision_name"`
	Office        string `json:"office_name" db:"office_name"`
}

// type EmployersResponse struct {
// 	Id           int    `json:"id" db:"id"`
// 	Name         string `json:"name" db:"name"`
// 	Attributes   map[string]string `json:"attributes" db:"extra_info"`
// 	Position     string `json:"position" db:"position"`
// 	DepartmentId int    `json:"department_id" db:"department_id"`
// 	ManagerId    int    `json:"manager_id" db:"manager_id"`
// 	Children     []EmployersResponse `json:"children"`
// }

// type Employers struct {
// 	Id           int    `json:"id" db:"id"`
// 	Name         string `json:"name" db:"name"`
// 	Attributes   string `json:"attributes" db:"extra_info"`
// 	Position     string `json:"position" db:"position"`
// 	DepartmentId int    `json:"department_id" db:"department_id"`
// 	ManagerId    int    `json:"manager_id" db:"manager_id"`
// }
