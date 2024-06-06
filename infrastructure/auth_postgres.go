package infrastructure

import (
	"Stepuha.net/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const DefaultMoneyCount = 0
const DefaultIsModerator = false
const DefaultIsFrozen = false

type AuthPostgres struct {
	Authorization
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (repos *AuthPostgres) AddUser(user entities.User) (int, error) {
	var id int
	user.Money = DefaultMoneyCount
	user.IsModerator = DefaultIsModerator
	user.IsFrozen = DefaultIsFrozen
	query := fmt.Sprintf("INSERT INTO %s (username, name, surname, teacher, money, password_hash) values ($1, $2, $3, $4, $5, $6) RETURNING id", UsersTable)
	row := repos.db.QueryRow(query, user.Username, user.Name, user.Surname, user.Teacher, user.Money, user.Password)
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

func (repos *AuthPostgres) CheckIfFrozen(userId int) (bool, error) {
	query := fmt.Sprintf("SELECT gd.is_frozen FROM %s gd WHERE gd.id = $1", UsersTable)
	var isFrozen bool
	err := repos.db.Get(&isFrozen, query, userId)
	print(isFrozen)
	if err != nil {
		return false, err
	}
	return isFrozen, nil
}
