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

type Good interface {
	Create(userId int, good entities.Good) (int, error)
	GetAll(userId int) ([]entities.Good, error)
	GetById(userId int, goodId int) (entities.Good, error)
	Delete(userId int, goodId int) error
	Update(userId int, goodId int, input entities.UpdateGoodInput) error
}

type Service struct {
	Authorization
	Good
}

func NewService(repository *infrastructure.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Good:          NewGoodService(repository.Good),
	}
}
