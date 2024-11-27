package repository

import (
	"rest_api_learn/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{
		db: db,
	}
}


func (db *AuthPostgres)	CreateUser(models.User) (int,error) {
	return 1, nil
}
