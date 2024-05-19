package infrastructure

import (
	"Stepuha.net/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (repos *UserPostgres) GetUserById(userId int) (entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT * FROM %s gd WHERE gd.id=$1", UsersTable)
	err := repos.db.Get(&user, query, userId)

	return user, err
}
