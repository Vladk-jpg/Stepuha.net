package infrastructure

import (
	"Stepuha.net/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	AddUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
}

type Good interface {
	Create(userId int, good entities.Good) (int, error)
	GetAll(userId int) ([]entities.Good, error)
	GetById(userId int, goodId int) (entities.Good, error)
	Delete(userId int, goodId int) error
	Update(userId int, goodId int, input entities.UpdateGoodInput) error
	CheckBelonging(userId int, goodId int) error
}

type Repository struct {
	Authorization
	Good
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Good:          NewGoodPostgres(db),
	}
}
