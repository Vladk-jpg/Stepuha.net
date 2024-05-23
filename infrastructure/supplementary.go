package infrastructure

import (
	"Stepuha.net/entities"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SupplementaryPostgres struct {
	db *sqlx.DB
}

func NewSupplementaryPostgres(db *sqlx.DB) *SupplementaryPostgres {
	return &SupplementaryPostgres{db: db}
}

func (repos *SupplementaryPostgres) GetRandomGoods(userId int) ([]entities.Good, error) {
	var goods []entities.Good
	query := fmt.Sprintf("SELECT gd.id, gd.name, gd.price, gd.picture, gd.description FROM %s gd ORDER BY RANDOM() LIMIT 15", GoodsTable)
	err := repos.db.Select(&goods, query)

	return goods, err
}

func (repos *SupplementaryPostgres) TransferMoney(senderId int, receiverId int, amount float64) error {
	tx, err := repos.db.Begin()
	if err != nil {
		return err
	}

	sendMoneyQuery := fmt.Sprintf("UPDATE %s SET money = money - $1 WHERE id = $2 AND money > 0", UsersTable)
	rows, err := tx.Exec(sendMoneyQuery, amount, senderId)

	if rows == nil {
		return errors.New("couldn't send money from user")
	}
	affectedRow, err := rows.RowsAffected()
	if affectedRow == 0 {
		transactionErr := tx.Rollback()
		if transactionErr != nil {
			return transactionErr
		}
		return errors.New("couldn't send money from user")
	}

	if err != nil {
		transactionErr := tx.Rollback()
		if transactionErr != nil {
			return transactionErr
		}
		return err
	}

	receiveMoneyQuery := fmt.Sprintf("UPDATE %s SET money = money + $1 WHERE id = $2", UsersTable)
	rows, err = tx.Exec(receiveMoneyQuery, amount, receiverId)

	if rows == nil {
		return errors.New("couldn't find the receiver")
	}
	affectedRow, err = rows.RowsAffected()
	if affectedRow == 0 {
		transactionErr := tx.Rollback()
		if transactionErr != nil {
			return transactionErr
		}
		return errors.New("couldn't send money from user")
	}

	if err != nil {
		transactionErr := tx.Rollback()
		if transactionErr != nil {
			return transactionErr
		}
		return err
	}

	return tx.Commit()
}
