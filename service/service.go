package service

import (
	"Stepuha.net/entities"
	"Stepuha.net/infrastructure"
)

type Authorization interface {
	AddUser(user entities.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repository *infrastructure.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
