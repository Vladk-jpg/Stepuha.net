package infrastructure

import (
	"Stepuha.net/entities"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	SSLMode  string
}

type Authorization interface {
	AddUser(user entities.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository() *Repository {
	return &Repository{}
}
