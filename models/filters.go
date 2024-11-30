package models


type PositionModel struct {
	Id int `json:"id" db:"id"`
	Position string `json:"position" db:"position"`
}


type FiltersResponse struct {
	Position map[string]string `json:"position"`
}