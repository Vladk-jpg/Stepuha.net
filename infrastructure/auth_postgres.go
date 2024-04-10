package infrastructure

import (
	"Stepuha.net/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	Authorization
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (repos *AuthPostgres) AddUser(user entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, teacher, password_hash) values ($1, $2, $3, $4), RETURNING id", UsersTable)
	row := repos.db.QueryRow(query, user.Name, user.Surname, user.Teacher, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
