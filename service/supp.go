package service

import (
	"Stepuha.net/entities"
	"Stepuha.net/infrastructure"
)

type SupplementaryService struct {
	repo infrastructure.Supplementary
}

func NewSupplementaryService(repo infrastructure.Supplementary) *SupplementaryService {
	return &SupplementaryService{repo: repo}
}

func (serv *SupplementaryService) GetRandomGoods(userId int) ([]entities.Good, error) {
	return serv.repo.GetRandomGoods(userId)
}

func (serv *SupplementaryService) TransferMoney(senderId int, receiverId int, amount float64) error {
	return serv.repo.TransferMoney(senderId, receiverId, amount)
}