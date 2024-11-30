package service

import (
	"rest_api_learn/models"
	"rest_api_learn/pgk/repository"
)

type FilterService struct {
	repos repository.Filters
}

func NewFilterService (repos repository.Filters) *FilterService {
	return &FilterService{
		repos: repos,
	}
}

func (s *FilterService) ReadAllFiltersDate() (models.FiltersResponse,error) {
	return s.repos.ReadAllFiltersDate()
}