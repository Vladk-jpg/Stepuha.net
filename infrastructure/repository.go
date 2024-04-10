package infrastructure

import (
	"Stepuha.net/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	AddUser(user entities.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
