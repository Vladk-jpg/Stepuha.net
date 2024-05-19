package service

import (
	"Stepuha.net/entities"
	"Stepuha.net/infrastructure"
)

type UserService struct {
	repo infrastructure.User
}

func (serv *UserService) GetUserById(userId int) (entities.User, error) {
	return serv.repo.GetUserById(userId)
}

func NewUserService(repo infrastructure.User) *UserService {
	return &UserService{repo: repo}
}
