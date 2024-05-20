package infrastructure

import (
	"Stepuha.net/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (repos *UserPostgres) UpdateUser(userId int, input entities.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}

	if input.Teacher != nil {
		setValues = append(setValues, fmt.Sprintf("teacher=$%d", argId))
		args = append(args, *input.Teacher)
		argId++
	}

	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password_hash=$%d", argId))
		args = append(args, *input.Password)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s gd SET %s WHERE gd.id = $%d",
		UsersTable, setQuery, argId)
	args = append(args, userId)

	_, err := repos.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
