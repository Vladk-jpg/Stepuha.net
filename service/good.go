package service

import (
	"Stepuha.net/entities"
	"Stepuha.net/infrastructure"
)

type GoodService struct {
	repo infrastructure.Good
}

func NewGoodService(repo infrastructure.Good) *GoodService {
	return &GoodService{repo: repo}
}

func (serv *GoodService) Create(userId int, good entities.Good) (int, error) {
	return serv.repo.Create(userId, good)
}

func (serv *GoodService) GetAll(userId int) ([]entities.Good, error) {
	return serv.repo.GetAll(userId)
}

func (serv *GoodService) GetGoodById(userId int, goodId int) (entities.Good, error) {
	return serv.repo.GetGoodById(userId, goodId)
}

func (serv *GoodService) Delete(userId int, goodId int) error {
	return serv.repo.Delete(userId, goodId)
}

func (serv *GoodService) Update(userId int, goodId int, input entities.UpdateGoodInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	err := serv.repo.Update(userId, goodId, input)
	return err
}

func (serv *GoodService) Buy(userId int, goodId int) error {
	return serv.repo.Buy(userId, goodId)
}
