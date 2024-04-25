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
	query := fmt.Sprintf("INSERT INTO %s (username, name, surname, teacher, password_hash) values ($1, $2, $3, $4, $5) RETURNING id", UsersTable)
	row := repos.db.QueryRow(query, user.Username, user.Name, user.Surname, user.Teacher, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repos *AuthPostgres) GetUser(username, password string) (entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", UsersTable)
	err := repos.db.Get(&user, query, username, password)

	return user, err
}
