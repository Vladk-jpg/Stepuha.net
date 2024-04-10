package service

import (
	"Stepuha.net/entities"
)

type Authorization interface {
	AddUser(user entities.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService() *Service {
	return &Service{}
}
