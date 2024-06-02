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
	GetGoodById(userId int, goodId int) (entities.Good, error)
	Delete(userId int, goodId int) error
	Update(userId int, goodId int, input entities.UpdateGoodInput) error
	Buy(userId int, goodId int) error
	CheckBelonging(userId int, goodId int) error
}

type User interface {
	GetUserById(userId int) (entities.User, error)
	GetYourUser(userId int) (entities.YourUser, error)
	UpdateUser(userId int, input entities.UpdateUserInput) error
}

type Supplementary interface {
	GetRandomGoods(userId int) ([]entities.Good, error)
	TransferMoney(senderId int, receiver int, amount int) error
	CheckIfModer(userId int) (bool, error)
	CheckIfFrozen(userId int) (bool, error)
	GetOwner(goodId int) (int, error)
}

type Moder interface {
	FreezeUser(userId int) error
}

type Service struct {
	Authorization
	Good
	User
	Moder
	Supplementary
}

func NewService(repository *infrastructure.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Good:          NewGoodService(repository.Good),
		User:          NewUserService(repository.User),
		Supplementary: NewSupplementaryService(repository.Supplementary),
		Moder:         NewModerService(repository.Moder),
	}
}
