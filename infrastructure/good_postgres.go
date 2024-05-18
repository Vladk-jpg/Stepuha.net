package infrastructure

import (
	"Stepuha.net/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type GoodPostgres struct {
	db *sqlx.DB
}

func NewGoodPostgres(db *sqlx.DB) *GoodPostgres {
	return &GoodPostgres{db: db}
}

func (repos *GoodPostgres) Create(userId int, good entities.Good) (int, error) {
	transaction, err := repos.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createGoodQuery := fmt.Sprintf("INSERT INTO %s (name, price, picture, description) VALUES ($1, $2, $3, $4) RETURNING id", GoodsTable)
	row := transaction.QueryRow(createGoodQuery, good.Name, good.Price, good.Picture, good.Description)
	if err = row.Scan(&id); err != nil {
		rollbackErr := transaction.Rollback()
		if rollbackErr != nil {
			log.Fatalf("Unable to rollback the insertion into " + GoodsTable)
		}
		return 0, err
	}

	createUsersGoodQuery := fmt.Sprintf("INSERT INTO %s (user_id, good_id) VALUES ($1, $2)", UsersGoodsTable)
	_, err = transaction.Exec(createUsersGoodQuery, userId, id)
	if err != nil {
		transactionErr := transaction.Rollback()
		if transactionErr != nil {
			log.Fatalf("Unable to rollback the insertion into " + UsersGoodsTable)
		}
		return 0, err
	}
	return id, transaction.Commit()
}

func (repos *GoodPostgres) GetAll(userId int) ([]entities.Good, error) {
	var goods []entities.Good
	query := fmt.Sprintf("SELECT gd.id, gd.name, gd.price, gd.picture, gd.description FROM %s gd INNER JOIN %s ugd on gd.id = ugd.good_id"+
		" WHERE ugd.user_id = $1", GoodsTable, UsersGoodsTable)
	err := repos.db.Select(&goods, query, userId)

	return goods, err
}

func (repos *GoodPostgres) GetById(userId int, goodId int) (entities.Good, error) {
	var good entities.Good
	query := fmt.Sprintf("SELECT gd.id, gd.name, gd.price, gd.picture, gd.description FROM %s gd INNER JOIN %s ugd on gd.id = ugd.good_id"+
		" WHERE ugd.user_id = $1 AND ugd.good_id = $2", GoodsTable, UsersGoodsTable)
	err := repos.db.Get(&good, query, userId, goodId)

	return good, err
}

func (repos *GoodPostgres) Delete(userId int, goodId int) error {
	query := fmt.Sprintf("DELETE FROM %s gd USING %s ugd WHERE gd.id = ugd.goodid AND ugd.user_id=$1 AND ugd.good_id=$2",
		GoodsTable, UsersTable)
	_, err := repos.db.Exec(query, userId, goodId)
	return err
}

func (repos *GoodPostgres) Update(userId int, goodId int, input entities.UpdateGoodInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.Picture != nil {
		setValues = append(setValues, fmt.Sprintf("picture=$%d", argId))
		args = append(args, *input.Picture)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s gd SET %s FROM %s ugd WHERE gd.id = ugd.good_id AND ugd.user_id=$%d AND ugd.good_id=$%d",
		GoodsTable, setQuery, UsersGoodsTable, argId, argId+1)
	args = append(args, userId, goodId)

	_, err := repos.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
