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

func (serv *UserService) GetYourUser(userId int) (entities.User, error) {
	return serv.repo.GetYourUser(userId)
}

func (serv *UserService) UpdateUser(userId int, input entities.UpdateUserInput) error {
	return serv.repo.UpdateUser(userId, input)
}

func NewUserService(repo infrastructure.User) *UserService {
	return &UserService{repo: repo}
}
