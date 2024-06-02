package service

import (
	"Stepuha.net/infrastructure"
)

type ModerService struct {
	repo infrastructure.Moder
}

func NewModerService(repo infrastructure.Moder) *ModerService {
	return &ModerService{repo: repo}
}

func (serv *ModerService) FreezeUser(userId int) error {
	return serv.repo.FreezeUser(userId)
}

func (serv *ModerService) UnfreezeUser(userId int) error {
	return serv.repo.UnfreezeUser(userId)
}
