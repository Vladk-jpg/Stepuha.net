package infrastructure

import (
	"Stepuha.net/entities"
	"fmt"
)

func (repos *GoodPostgres) CheckBelonging(userId int, goodId int) error {
	var good entities.Good
	query := fmt.Sprintf("SELECT gd.id FROM %s gd INNER JOIN %s ugd on gd.id = ugd.good_id"+
		" WHERE ugd.user_id = $1 AND ugd.good_id = $2", GoodsTable, UsersGoodsTable)
	err := repos.db.Get(&good, query, userId, goodId)
	if err != nil {
		return err
	}

	return nil
}
