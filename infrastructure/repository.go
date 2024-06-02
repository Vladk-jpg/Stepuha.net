package infrastructure

import (
	"Stepuha.net/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	AddUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
	CheckIfFrozen(userId int) (bool, error)
}

type Good interface {
	Create(userId int, good entities.Good) (int, error)
	GetAll(userId int) ([]entities.Good, error)
	GetGoodById(userId int, goodId int) (entities.Good, error)
	Delete(userId int, goodId int) error
	Update(userId int, goodId int, input entities.UpdateGoodInput) error
	CheckBelonging(userId int, goodId int) error
	Buy(userId int, goodId int) error
}

type User interface {
	GetUserById(userId int) (entities.User, error)
	GetYourUser(userId int) (entities.YourUser, error)
	UpdateUser(userId int, input entities.UpdateUserInput) error
}

type Supplementary interface {
	GetRandomGoods(userId int) ([]entities.Good, error)
	TransferMoney(senderId int, receiverId int, amount int) error
	CheckIfModer(userId int) (bool, error)
	GetOwner(goodId int) (int, error)
}

type Moder interface {
	FreezeUser(userId int) error
	UnfreezeUser(userId int) error
}

type Repository struct {
	Authorization
	Good
	User
	Moder
	Supplementary
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Good:          NewGoodPostgres(db),
		User:          NewUserPostgres(db),
		Supplementary: NewSupplementaryPostgres(db),
		Moder:         NewModerPostgres(db),
	}
}
